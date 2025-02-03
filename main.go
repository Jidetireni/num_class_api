package main

import (
	"num_class_api/handlers"
	"num_class_api/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	handlers.RegisterRoutes(server)
	server.Run(":8080")

}

func Funfact(n int) (string, error) {
	funfact, err := utils.GetFuncFact(n)
	if err != nil {
		return "", err
	}

	return funfact, nil
}
