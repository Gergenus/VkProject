package handlers

import (
	"net/http"
	"strings"

	"github.com/Gergenus/VkProject/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var allowePhotoTypes = map[string]bool{
	"jpg": true,
	"png": true,
}

func (u *UserHandlers) CreatePost(c echo.Context) error {
	var data models.Post

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}

	if len(data.PostText) > 2500 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "contents should be less than 2500 chars",
		})
	}

	if len(data.Subject) > 100 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "subject should be less than 100 chars",
		})
	}

	if data.Price < 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "subject cannot be negative",
		})
	}

	if !linkValidation(data.ImageAddress) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid photo type",
		})
	}

	resp, err := http.Head(data.ImageAddress)
	if err != nil {
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"error": "failed to get file size",
		})
	}
	if resp.ContentLength > int64(5242880) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "content size should be less than 5mb",
		})
	}
	defer resp.Body.Close()

	uid, ok := c.Get("uid").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}
	data.UserID = uuid.MustParse(uid)

	id, err := u.postSrv.CreatePost(c.Request().Context(), data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":      id,
		"subject": data.Subject,
	})
}

func linkValidation(imageAddress string) bool {
	splitLink := strings.Split(imageAddress, ".")
	format := splitLink[len(splitLink)-1]
	_, ok := allowePhotoTypes[format]
	return ok
}
