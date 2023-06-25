package entryPoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, err *EndpointError, code int) {
	c.JSON(code, &EndpointResponse{
		Success: false,
		Error:   err,
	})
}

func SendBadGateAway(c *gin.Context, message, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadGateway,
			Message:   message,
			Origin:    origin,
		},
	})
}

func SendInvalidTokenError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrInvalidToken,
			Origin:    origin,
		},
	})
}

func SendInternalServerError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusInternalServerError,
			Message:   ErrInternalServerError,
			Origin:    origin,
		},
	})
}

func SendInvalidUserIdError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrInvalidUserId,
			Origin:    origin,
		},
	})
}

func SendUserNotFoundError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusNotFound,
			Message:   ErrPackNotFound,
			Origin:    origin,
		},
	})
}

func SendRestoredOnlyError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrRestoredOnly,
			Origin:    origin,
		},
	})
}

func SendUserNotRegisteredError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusMethodNotAllowed,
			Message:   ErrUserNotRegistered,
			Origin:    origin,
		},
	})
}

func SendNoDataError(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrNoData,
			Origin:    origin,
		},
	})
}

func SendPermissionDenied(c *gin.Context, origin string) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: false,
		Error: &EndpointError{
			ErrorCode: http.StatusBadRequest,
			Message:   ErrPermissionDenied,
			Origin:    origin,
		},
	})
}

func SendResult(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, &EndpointResponse{
		Success: true,
		Result:  result,
	})
}
