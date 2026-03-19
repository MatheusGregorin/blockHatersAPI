package main

import (
	"fmt"
	"log"
	"myMarket/handler"
	"myMarket/internal/database"
	"myMarket/internal/middleware"
	"myMarket/internal/repository"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func main() {

	// carregando .ENV
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar o arquivo .env")
	}

	// Rate limiter configuration
	rate := limiter.Rate{
		Period: 5 * time.Second, // Período de tempo para o qual a taxa é calculada
		Limit:  3,               // Qtd requisições permitidas por período
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	middlewareRateLimit := mgin.NewMiddleware(instance)

	// Cors configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	// Database connection
	database.Connect()

	// Bin Repository // MYSQL
	UserRepository := repository.NewUserMysqlRepository()
	repositoryType := os.Getenv("REPOSITORY")
	if repositoryType == "postgres" {
		// UserRepository = repository.NewUserRepositoryPostgres() // Implemente esta função para retornar um repositório PostgreSQL
	}
	userHandler := handler.NewUserHandler(UserRepository)

	// Init Gin
	r := gin.Default()

	// Registro das configurações
	r.Use(middlewareRateLimit)
	r.Use(cors.New(config))

	// Rota de autenticação
	r.POST("/login", userHandler.Login)

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

	fmt.Println("Servidor rodando na porta 8083")

	r.Run(":8083")
}
