package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors,omitempty"`
}

// Success Response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// Created Response
func CreatedResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusCreated, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// No Content Response
func NoContentResponse(c *gin.Context, code int) {
	c.JSON(http.StatusNoContent, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Bad Request Response
func ErrorBadRequestResponse(c *gin.Context, code int, errors interface{}) {
	var errMessages []string
	switch e := errors.(type) {
	case string:
		errMessages = []string{e}
	case []string:
		errMessages = e
	default:
		errMessages = []string{"Invalid request"}
	}
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
		Errors:  errMessages,
	})
}

// Error Unauthorized Response
func ErrorUnauthorizedResponse(c *gin.Context, code int) {
	c.JSON(http.StatusUnauthorized, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Forbidden Response
func ErrorForbiddenResponse(c *gin.Context, code int) {
	c.JSON(http.StatusForbidden, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Not Found Response
func ErrorNotFoundResponse(c *gin.Context, code int) {
	c.JSON(http.StatusNotFound, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Method Not Allowed Response
func ErrorMethodNotAllowedResponse(c *gin.Context, code int) {
	c.JSON(http.StatusMethodNotAllowed, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Conflict Response
func ErrorConflictResponse(c *gin.Context, code int) {
	c.JSON(http.StatusConflict, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Payload Too Large Response
func ErrorPayloadTooLargeResponse(c *gin.Context, code int) {
	c.JSON(http.StatusRequestEntityTooLarge, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error URI Too Long Response
func ErrorURITooLongResponse(c *gin.Context, code int) {
	c.JSON(http.StatusRequestURITooLong, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Unsupported Media Type Response
func ErrorUnsupportedMediaTypeResponse(c *gin.Context, code int) {
	c.JSON(http.StatusUnsupportedMediaType, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Too Many Requests Response
func ErrorTooManyRequestsResponse(c *gin.Context, code int) {
	c.JSON(http.StatusTooManyRequests, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Internal Server Response
func ErrorInternalServerResponse(c *gin.Context, code int) {
	c.JSON(http.StatusInternalServerError, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Not Implemented Response
func ErrorNotImplementedResponse(c *gin.Context, code int) {
	c.JSON(http.StatusNotImplemented, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Bad Gateway Response
func ErrorBadGatewayResponse(c *gin.Context, code int) {
	c.JSON(http.StatusBadGateway, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Service Unavailable Response
func ErrorServiceUnavailableResponse(c *gin.Context, code int) {
	c.JSON(http.StatusServiceUnavailable, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// Error Gateway Timeout Response
func ErrorGatewayTimeoutResponse(c *gin.Context, code int) {
	c.JSON(http.StatusGatewayTimeout, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
