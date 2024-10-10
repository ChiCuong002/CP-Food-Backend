package apierror

import (
	"fmt"
	"food-recipes-backend/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type APIError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg"`
}
type CustomError struct {
	StatusCode int   `json:"code"`
	Message    string `json:"msg"`
}
func (e *CustomError) Error() string {
	return e.Message
}
func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}
func NewCustomError(statusCode int, message string) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
	}
}
func NewAPIError(statusCode int, message string) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}
func InvalidRequestData() APIError {
	return APIError{
		StatusCode: 400,
		Message:    "invalid request data",
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, "invalid JSON request data")
}


type APIFunc func(c *gin.Context) error

func Make(h APIFunc) gin.HandlerFunc {
    return func(c *gin.Context) {
        if err := h(c); err != nil {
            if apiErr, ok := err.(APIError); ok {
                writeJSON(c, apiErr.StatusCode, apiErr)
            } else{
                errResp := map[string]any{
                    "code": http.StatusInternalServerError,
                    "msg":  "internal server error",
                }
                writeJSON(c, http.StatusInternalServerError, errResp)
				global.Logger.Error("internal server error", zap.Error(err))
            }
        }
    }
}

func writeJSON(c *gin.Context, statusCode int, data interface{}) {
    c.JSON(statusCode, data)
}
func ErrorHandler(c *gin.Context, err APIError) {
    c.JSON(err.StatusCode, err)
}