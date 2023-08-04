package main

import (
	config "FirstAssignment/config"
	"FirstAssignment/database"
	"FirstAssignment/service/domain/user/delivery"
	"FirstAssignment/service/domain/user/usecase"
	repo2 "FirstAssignment/service/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	schema := config.NewSchema()
	db, err := database.Open(schema.SQL.DataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	repo := repo2.NewUserRepo(db)
	userUseCase := usecase.NewUserUseCase(repo)
	userHandler := delivery.NewUserHandler(userUseCase)

	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
	router := server.Group("/")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	userHandler.UserAPIRoute(router)

	log.Fatal(server.Run(":3000"))
}
