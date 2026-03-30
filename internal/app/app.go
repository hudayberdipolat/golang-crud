package app

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func NewApp(dependencies *Dependency) (httpServer *fiber.App) {

	httpServer = fiber.New(fiber.Config{
		AppName:      dependencies.Config.ServerConfig.AppName,
		BodyLimit:    35 * 1024 * 1024,
		ServerHeader: dependencies.Config.ServerConfig.AppHeader,
		WriteTimeout: time.Minute * 3,
		ReadTimeout:  time.Minute * 3,
	})
	httpServer.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept, API-KEY, Language, Platform, Authorization, Access-Control-Allow-Origin"},
		AllowMethods: []string{"GET,POST,HEAD,PUT,DELETE,PATCH"},
	}))
	return httpServer
}
