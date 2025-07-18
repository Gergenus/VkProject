package middlew

import (
	"net/http"
	"os"

	"github.com/Gergenus/VkProject/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func TestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		coookie, err := c.Cookie("AccessToken")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "no auth token",
			})
		}
		if coookie == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "no auth token",
			})
		}
		token := coookie.Value
		uid, login, err := jwt.ParseToken(token, os.Getenv("JWT_SECRET"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "inernal server error",
			})
		}

		c.Set("uid", uid)
		c.Set("login", login)

		return next(c)
	}
}
