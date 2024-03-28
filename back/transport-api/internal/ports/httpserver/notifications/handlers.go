package notifications

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transport-api/internal/app"
	"transport-api/internal/model"
	"transport-api/internal/ports/httpserver/middleware"
)

func GetNotifications(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		onlyNotViewedStr, _ := c.GetQuery("only_not_viewed")
		onlyNotViewed := onlyNotViewedStr == "true"

		notifications, err := a.GetNotifications(c, companyId, uint(limit), uint(offset), onlyNotViewed)
		switch {
		case err == nil:
			c.JSON(http.StatusOK, notificationListResponse{
				Data: notificationsToNotificationsDataList(notifications),
				Err:  nil,
			})
		case errors.Is(err, model.ErrNotificationsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsUnknown))
		}
	}
}

func GetNotification(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		notification, err := a.GetNotification(c, companyId, id)

		switch {
		case err == nil:
			data := notificationToNotificationData(notification)
			c.JSON(http.StatusOK, notificationResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrNotificationNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrNotificationNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrNotificationsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsUnknown))
		}
	}
}

func SubmitClosedLead(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, companyId, ok := middleware.GetAuthData(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(model.ErrUnauthorized))
			return
		}
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		submitStr, _ := c.GetQuery("submit")
		submit := submitStr == "true"

		err = a.SubmitClosedLead(c, companyId, id, submit)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, errorResponse(nil))
		case errors.Is(err, model.ErrNotificationNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrNotificationNotExists))
		case errors.Is(err, model.ErrPermissionDenied):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrPermissionDenied))
		case errors.Is(err, model.ErrInvalidInput):
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
		case errors.Is(err, model.ErrNotificationAnswered):
			c.AbortWithStatusJSON(http.StatusConflict, errorResponse(model.ErrNotificationAnswered))
		case errors.Is(err, model.ErrNotificationsError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrNotificationsUnknown))
		}
	}
}
