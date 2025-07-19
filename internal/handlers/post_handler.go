package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Gergenus/VkProject/internal/models"
	"github.com/Gergenus/VkProject/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (u *UserHandlers) CreatePost(c echo.Context) error {
	var data models.ProductPost
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid payload",
		})
	}
	uid, ok := c.Get("uid").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}
	data.UserID = uuid.MustParse(uid)

	id, err := u.postSrv.CreatePost(c.Request().Context(), data)
	if err != nil {
		if errors.Is(err, service.ErrHeadRequestFailed) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		if errors.Is(err, service.ErrIncorrectContents) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		if errors.Is(err, service.ErrIncorrectImageAddress) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		if errors.Is(err, service.ErrIncorrectImageSize) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		if errors.Is(err, service.ErrIncorrectPrice) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		if errors.Is(err, service.ErrIncorrectSubject) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":      id,
		"subject": data.Subject,
	})
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
