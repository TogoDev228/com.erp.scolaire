package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// InitDB initialise la connexion à la base de données
func InitDB() *gorm.DB {
	once.Do(func() {
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")

		if user == "" || host == "" || port == "" || name == "" {
			log.Fatal("Les variables d'environnement de la base de données ne sont pas toutes définies.")
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, name)

		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Erreur de connexion à la base de données : %v", err)
		}

		fmt.Println("Connexion à la base de données réussie.")
	})

	return DB
}
