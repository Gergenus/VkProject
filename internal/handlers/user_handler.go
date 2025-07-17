package handlers

import (
	"errors"
	"net/http"

	"github.com/Gergenus/VkProject/internal/models"
	"github.com/Gergenus/VkProject/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	srv service.UserServiceInterface
}

func NewUserHandler(srv service.UserServiceInterface) *UserHandlers {
	return &UserHandlers{
		srv: srv,
	}
}

func (u *UserHandlers) SignUp(c echo.Context) error {
	var user models.RegisterRequest

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}
	if len(user.Login) < 3 || len(user.Login) > 50 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "login must be 3-50 characters",
		})
	}
	if len(user.Password) < 8 || len(user.Password) > 50 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "password must be 8-50 characters",
		})
	}

	addedUser, err := u.srv.RegisterNewUser(c.Request().Context(), user.Login, user.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserExists) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "user already exists",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"uuid":  addedUser.ID.String(),
		"login": addedUser.Login,
	})
}
