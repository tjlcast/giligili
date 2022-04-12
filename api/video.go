package api

import (
	"giligili/serializer"

	"github.com/gin-gonic/gin"
)

// CreateVideo
func CreateVideo(c *gin.Context) {
	// service := service.Service{}

	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Logout",
	})
}
