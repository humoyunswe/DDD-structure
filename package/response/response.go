package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	RequestID string    `json:"request_id"`
	Timestamp time.Time `json:"timestamp"`
}

type APIResponse struct {
	Success bool         `json:"success"`
	Data    interface{}  `json:"data,omitempty"`
	Total   *int64       `json:"total,omitempty"`
	Limit   *int         `json:"limit,omitempty"`
	Offset  *int         `json:"offset,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
	Meta    Meta         `json:"meta"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
		Meta:    newMeta(c),
	})
}

func SuccessPaginated(c *gin.Context, data interface{}, total int64, limit, offset int) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
		Total:   &total,
		Limit:   &limit,
		Offset:  &offset,
		Meta:    newMeta(c),
	})
}

func Fail(c *gin.Context, statusCode int, code, message string) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
		},
		Meta: newMeta(c),
	})
}

func newMeta(c *gin.Context) Meta {
	return Meta{
		RequestID: generateRequestID(c),
		Timestamp: time.Now().UTC(),
	}
}

func generateRequestID(c *gin.Context) string {
	id := c.GetString("request_id")
	if id == "" {
		id = uuid.NewString()
		c.Set("request_id", id)
	}
	return id
}
