package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ReturnFail return without error message
func ReturnFail(r *gin.Context) {
	r.JSON(200, gin.H{
		"code":    -1,
		"message": "fail",
	})
}

// ReturnError return with error message
func ReturnError(r *gin.Context, err error) {
	r.JSON(200, gin.H{
		"code":    -1,
		"message": fmt.Sprintf("error: %v", err),
	})
}

// ReturnOK return without data
func ReturnOK(r *gin.Context) {
	r.JSON(200, gin.H{
		"code":    0,
		"message": "ok",
	})
}

// ReturnSuccess return with data
func ReturnSuccess(r *gin.Context, data interface{}) {
	r.JSON(200, gin.H{
		"code":   0,
		"messae": "success",
		"data":   data,
	})
}
