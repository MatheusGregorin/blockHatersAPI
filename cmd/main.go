package main

import (
	"fmt"
	"log"
	"myMarket/handler"
	"myMarket/internal/database"
	"myMarket/internal/middleware"
	"myMarket/internal/repository"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar o arquivo .env")
	}

	database.Connect()

	// MYSQL - DEFAULT
	UserMysqlRepository := repository.NewUserMysqlRepository()

	repositoryType := os.Getenv("REPOSITORY")
	if repositoryType == "postgres" {
		// UserMysqlRepository = repository.NewUserMysqlRepositoryPostgres() // Implemente esta função para retornar um repositório PostgreSQL
	}

	// Bindando o repositório ao handler
	userHandler := handler.NewUserHandler(UserMysqlRepository)

	r := gin.Default()
	r.POST("/login", userHandler.Login)
	r.POST("/add", userHandler.Register)

	// // Grupo de rotas protegidas
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.POST("/register", userHandler.Register)
		v1.GET("/user/:id", userHandler.GetUserByID)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Rota não encontrada"})
	})

	fmt.Println("\nServidor rodando na porta 8083\n")
	r.Run(":8083")
}
