package main

import (
	"context"
	"fmt"
	"gotbittask/internal/app/task/delivery"
	"gotbittask/internal/app/task/repository"
	"gotbittask/internal/app/task/usecase"
	"log"

	_ "gotbittask/docs"

	"github.com/fasthttp/router"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/valyala/fasthttp"
)

// @title           GotBit API
// @version         1.0
// @description     API server for tasks list application.

// @host      localhost:8081
// @BasePath  /api/v1

func main() {
	router := router.New()

	dbpool, err := pgxpool.Connect(context.Background(),
		"host=localhost port=5432 user=ilyagu dbname=gotbit password=password sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Success connection")
	defer dbpool.Close()

	// repositories
	taskRepo := repository.NewTaskRepo(dbpool)

	// usecases
	taskUsecase := usecase.NewTaskUsecase(taskRepo)

	// delivery
	delivery.NewTaskHandler(router, taskUsecase)

	// swagger
	fmt.Print("\nDocumentaion on http://localhost:8080/swagger/index.html\n\n")

	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

	err = fasthttp.ListenAndServe(":8081", router.Handler)
	log.Fatal(err)
}
