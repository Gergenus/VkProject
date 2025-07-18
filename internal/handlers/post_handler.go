package handlers

import (
	"net/http"
	"strconv"
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "failed to get file size",
		})
	}
	if resp.ContentLength > int64(5*1024*1024) {
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

func (p *UserHandlers) Posts(c echo.Context) error {
	pageString := c.QueryParam("page")
	if pageString == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}

	pageSizeString := c.QueryParam("page_size")
	if pageSizeString == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}

	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}

	sortBy := c.QueryParam("sort_by")
	if sortBy != "price" && sortBy != "created_at" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}

	sortDir := c.QueryParam("sort_dir")
	if sortDir != "asc" && sortDir != "desc" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}

	userId, ok := c.Get("uid").(string)
	if !ok {
		userId = ""
	}
	var minPrice float64
	minPriceString := c.QueryParam("min_price")
	if minPriceString != "" {
		minPrice, err = strconv.ParseFloat(minPriceString, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid payload",
			})
		}
	}
	var maxPrice float64
	maxPriceString := c.QueryParam("max_price")
	if maxPriceString != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceString, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid payload",
			})
		}
	}

	posts, err := p.postSrv.Posts(c.Request().Context(), page, pageSize, userId, sortBy, sortDir, minPrice, maxPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}
	return c.JSON(http.StatusOK, posts)
}
