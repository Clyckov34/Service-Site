package errs

import (
	"net/http"
	"repair/pkg/errs"

	"github.com/gin-gonic/gin"
)

//NotFound обработка ошибки
func NotFound(c *gin.Context) {
	errs.HTTPStatusJson(c, http.StatusNotFound, "Not Found")
}
