package main

import (
	"gin-blog/internal/models"
	"gin-blog/internal/pkg/config"
	"gin-blog/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	err := config.StartSetting()
	if err != nil {
		log.Fatalf("configuration startup failure: %v", err)
	}

	err = models.InitDB()
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}
}

func main() {

	gin.SetMode(config.Server.RunMode)

	r := routes.InitRouter()

	r.Run(":" + config.Server.HttpPort)
}
