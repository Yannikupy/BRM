package httpserver

import (
	"errors"
	"github.com/gin-gonic/gin"
	"images/internal/app"
	"images/internal/model"
	"io"
	"net/http"
	"strconv"
)

func handleAddImage(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		defer func() {
			_ = file.Close()
		}()

		data, err := io.ReadAll(file)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
			return
		}

		id, err := a.AddImage(c, data)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, idResponse{
				Id:  &id,
				Err: nil,
			})
		case errors.Is(err, model.ErrImageTooBig):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrImageTooBig))
		case errors.Is(err, model.ErrWrongFormat):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrWrongFormat))
		case errors.Is(err, model.ErrDatabaseError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrDatabaseError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}
	}
}

func handleGetImage(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		img, err := a.GetImage(c, id)

		switch {
		case err == nil:
			c.Data(http.StatusOK, http.DetectContentType(img), img)
		case errors.Is(err, model.ErrImageNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrImageNotExists))
		case errors.Is(err, model.ErrDatabaseError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrDatabaseError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrServiceError))
		}
	}
}
