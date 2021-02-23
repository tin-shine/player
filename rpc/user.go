package rpc

import (
	"log"

	"github.com/game_marketing/player/proto/user/model"
	"github.com/game_marketing/player/proto/user/proto/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var c user.InfoClient

// GetUserInfo GetUserInfo from remote service
func GetUserInfo(r *gin.Context) (*model.UserModel, error) {
	log.Println("get.user.info.from.remote")

	log.Println("init.rpc")
	conn, err := grpc.Dial("localhost:2021", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("failed to connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c = user.NewInfoClient(conn)
	resp, err := c.GetInfo(r, &user.InfoRequest{
		Uid: 123,
	})
	if err != nil {
		log.Printf("failed to get info: %v", err)
		return nil, err
	}

	return &model.UserModel{
		UID:  resp.Uid,
		Name: resp.Name,
		Sex:  resp.Sex,
	}, nil
}
