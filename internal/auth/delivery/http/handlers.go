package http

import (
	"bytes"
	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/auth"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/internal/session"
	"github.com/AleksK1NG/api-mc/pkg/httpErrors"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

// Auth handlers
type handlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	sessUC session.UCSession
}

// Auth handlers constructor
func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, sessUC session.UCSession) auth.Handlers {
	return &handlers{cfg, authUC, sessUC}
}

// Register new user
func (h *handlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		user := &models.User{}

		if err := utils.ReadRequest(c, user); err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		createdUser, err := h.authUC.Register(ctx, user)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		sess, err := h.sessUC.CreateSession(ctx, &models.Session{
			UserID: createdUser.User.UserID,
		}, h.cfg.Session.Expire)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		c.SetCookie(utils.CreateSessionCookie(h.cfg, sess))

		return c.JSON(http.StatusCreated, createdUser)
	}
}

// Login user
func (h *handlers) Login() echo.HandlerFunc {
	// Login user, validate email and password input
	type Login struct {
		Email    string `json:"email" db:"email" validate:"omitempty,lte=60,email"`
		Password string `json:"password,omitempty" db:"password" validate:"required,gte=6"`
	}
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		login := &Login{}

		if err := utils.ReadRequest(c, login); err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		userWithToken, err := h.authUC.Login(ctx, &models.User{
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		sess, err := h.sessUC.CreateSession(ctx, &models.Session{
			UserID: userWithToken.User.UserID,
		}, h.cfg.Session.Expire)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		c.SetCookie(utils.CreateSessionCookie(h.cfg, sess))

		return c.JSON(http.StatusOK, userWithToken)
	}
}

// Logout user
func (h *handlers) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		cookie, err := c.Cookie("session-id")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError(err))
			}
			return c.JSON(http.StatusInternalServerError, httpErrors.NewInternalServerError(err))
		}

		if err := h.sessUC.DeleteByID(ctx, cookie.Value); err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		utils.DeleteSessionCookie(c, h.cfg.Session.Name)

		return c.NoContent(http.StatusOK)
	}
}

// Update existing user
func (h *handlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		user := &models.User{}
		user.UserID = uID

		if err := utils.ReadRequest(c, user); err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		updatedUser, err := h.authUC.Update(ctx, user)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		return c.JSON(http.StatusOK, updatedUser)
	}
}

// Get user by id
func (h *handlers) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		user, err := h.authUC.GetByID(ctx, uID)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		return c.JSON(http.StatusOK, user)
	}
}

// Delete user handler
func (h *handlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		if err := h.authUC.Delete(ctx, uID); err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		return c.NoContent(http.StatusOK)
	}
}

// Find users by name
func (h *handlers) FindByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		if c.QueryParam("name") == "" {
			return c.JSON(http.StatusBadRequest, httpErrors.NewBadRequestError("name is required"))
		}

		paginationQuery, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		response, err := h.authUC.FindByName(ctx, c.QueryParam("name"), paginationQuery)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		return c.JSON(http.StatusOK, response)
	}
}

// Gat all users with pagination page and size query params
func (h *handlers) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		paginationQuery, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		usersList, err := h.authUC.GetUsers(ctx, paginationQuery)
		if err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		return c.JSON(http.StatusOK, usersList)
	}
}

// Load current user from ctx with auth middleware
func (h *handlers) GetMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*models.User)
		if !ok {
			return utils.ErrResponseWithLog(c, httpErrors.NewUnauthorizedError(httpErrors.Unauthorized))
		}

		return c.JSON(http.StatusOK, user)
	}
}

// Upload user avatar
func (h *handlers) UploadAvatar() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, cancel := utils.GetCtxWithReqID(c)
		defer cancel()

		image, err := utils.ReadImage(c, "avatar")
		if err != nil {
			return httpErrors.NewInternalServerError(err)
		}

		file, err := image.Open()
		if err != nil {
			return httpErrors.NewInternalServerError(err)
		}
		defer file.Close()

		binaryImage := bytes.NewBuffer(nil)

		if _, err := io.Copy(binaryImage, file); err != nil {
			return httpErrors.NewInternalServerError(err)
		}

		if _, err := utils.CheckImageFileContentType(binaryImage.Bytes()); err != nil {
			return httpErrors.NewBadRequestError(err)
		}

		if err := h.authUC.UploadAvatar(ctx, image.Filename, binaryImage.Bytes()); err != nil {
			return utils.ErrResponseWithLog(c, err)
		}

		return c.NoContent(http.StatusOK)
	}
}
