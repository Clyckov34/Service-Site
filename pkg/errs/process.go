package errs

import "github.com/gin-gonic/gin"

//HTTPStatusJson вывод сообщения
func HTTPStatusJson(ctx *gin.Context, httpStatus int, message string) {
	ctx.JSON(httpStatus, gin.H{
		"status":  httpStatus,
		"message": message,
	})
}
