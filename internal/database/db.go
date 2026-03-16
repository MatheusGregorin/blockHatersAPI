package database

import (
	"fmt"
	"log"
	"myMarket/internal/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// String de conexão (DSN)
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Falha na conexão: %v\n", err)
		return
	}

	DB = database
	fmt.Println("\nConexão com o banco de dados estabelecida com sucesso!\n")

	err = DB.AutoMigrate(&models.User{}, &models.Merchant{}, &models.Product{})
	if err != nil {
		log.Fatal("Falha ao rodar a migração: ", err)
	}
	fmt.Println("Migração do banco de dados concluída com sucesso!\n")
}
