package main

import (
	"github.com/game_marketing/player/proto/user/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get_info", handler.GetUserInfo)
	r.Run(":8081")
}
