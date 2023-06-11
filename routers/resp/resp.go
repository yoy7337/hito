package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CodePrefix int

const (
	System CodePrefix = iota
)

func (cp CodePrefix) String() string {
	return []string{"00"}[cp]
}

const success = 0

// the error response with http status code, error code and message
type Error struct {
	HttpStatusCode int
	Code           string
	Message        string
}

func (apiError Error) Error() string {
	return apiError.Message
}

func JSON(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": success, "data": obj})
}

func Err(c *gin.Context, err Error) {
	c.Error(err)
}

func ErrMsg(c *gin.Context, err Error, msg string) {
	err.Message = err.Error() + ": " + msg
	c.Error(err)
}
