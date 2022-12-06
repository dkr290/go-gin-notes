package models

import (
	"fmt"

	"github.com/dkr290/go-devops/go-gin-notes/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repo = repository.NewRepo()

func ConnectDatabase() {

	// Repo.DbHost = os.Getenv("DATABASE_HOST")
	// Repo.DbPass = os.Getenv("DATABASE_PASS")
	// Repo.DbUser = os.Getenv("DATABASE_USER")
	// Repo.DbPort = os.Getenv("DATABASE_PORT")
	// Repo.DbName = os.Getenv("DB_NAME")

	Repo.DbHost = "172.18.159.11"
	Repo.DbPass = "Password123"
	Repo.DbUser = "postgres"
	Repo.DbPort = "5432"
	Repo.DbName = "notes"

	var err error
	fmt.Printf("Host: %s, NAME: %s, USER: %s, PORT:%s,PASS: %s", Repo.DbHost, Repo.DbName, Repo.DbUser, Repo.DbPort, Repo.DbPass)
	dsn := "host=" + Repo.DbHost + " " + "user=" + Repo.DbUser + " " + "password=" + Repo.DbPass + " " + "dbname=" + Repo.DbName + " " + "port=" + Repo.DbPort
	Repo.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Connected to Database sucess!")

}

func DbMigrate() {
	Repo.DB.AutoMigrate(&Note{})
	Repo.DB.AutoMigrate(&User{})
}
