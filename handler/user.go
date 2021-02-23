package handler

import (
	"log"

	"github.com/game_marketing/player/proto/user/rpc"
	"github.com/game_marketing/player/proto/user/utils"
	"github.com/gin-gonic/gin"
)

// GetUserInfo GetUserInfo
func GetUserInfo(r *gin.Context) {
	log.Println("get.user.info")

	user, err := rpc.GetUserInfo(r)
	if err != nil {
		log.Printf("failed to GetUserInfo: %v", err)
		utils.ReturnFail(r)
		return
	}

	log.Println("get.user.info.success")
	utils.ReturnSuccess(r, user)
}
