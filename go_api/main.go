package main

import (
	"HDYS-TTBYS/my-todo/ent"
	"HDYS-TTBYS/my-todo/ent/migrate"
	infra "HDYS-TTBYS/my-todo/infrastructure/repository"
	"HDYS-TTBYS/my-todo/interfaces/handler"
	"HDYS-TTBYS/my-todo/usecase"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	toBool, err := strconv.ParseBool(os.Getenv("DEV"))
	var sslmode string
	if err != nil {
		sslmode = ""
	}
	if toBool {
		sslmode = " sslmode=disable"
	} else {
		sslmode = ""
	}
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s"+sslmode,
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	client, err := ent.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	context := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(
		context,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(os.Getenv("ALLOWORIGINS"), " "),
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "UPDATE", "OPTIONS", "PATCH"},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())

	todoInfra := infra.NewTodoRepository(client, context)
	todoUsecase := usecase.NewTodoUseCase(todoInfra)
	todoHandler := handler.NewTodoHandler(todoUsecase)
	api := e.Group("/api")
	api.GET("/todos", todoHandler.FindMany)
	api.GET("/todo/:id", todoHandler.FindByID)
	api.DELETE("/todo/:id", todoHandler.Delete)
	api.PATCH("/todo/:id", todoHandler.Update)
	api.POST("/todo", todoHandler.Create)
	e.Logger.Fatal(e.Start(":8080"))
}
