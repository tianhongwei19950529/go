package middlewares

import (
	"gincode/api/lib/idgen"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestId string

func (r RequestId) String() string {
	return string(r)
}

func Access(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 8*1024*1024)
	requestIdNames := []string{"requestId", "requestID", "request_id", "REQUEST_ID", "REQUEST-ID"}
	var requestId RequestId
	for _, idName := range requestIdNames {
		if c.GetHeader(idName) != "" {
			requestId = RequestId(c.GetHeader(idName))
			break
		}
	}
	if requestId == "" {
		requestId = RequestId(idgen.GenId())
	}
	c.Set("requestId", requestId)
	c.Next()
}
