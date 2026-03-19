package database

import (
	"fmt"
	"log"
	"myMarket/internal/models"
	"os"

	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// String de conexão (DSN)
	dsn := os.Getenv("DATABASE_URL")
	if !strings.Contains(dsn, "parseTime=true") {
		if strings.Contains(dsn, "?") {
			dsn += "&parseTime=true"
		} else {
			dsn += "?parseTime=true"
		}
	}
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Printf("Falha na conexão: %v\n", err)
		return
	}

	DB = database
	fmt.Println("\nConexão com o banco de dados estabelecida com sucesso!")

	err = DB.AutoMigrate(&models.Merchant{}, &models.User{}, &models.Review{})
	if err != nil {
		log.Fatal("Falha ao rodar a migração: ", err)
	}
	fmt.Println("Migração do banco de dados concluída com sucesso!")

	var users []models.User
	DB.Find(&users)
	fmt.Println(users)
}
