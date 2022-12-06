package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/dkr290/go-gin-notes/app/config"
	"github.com/dkr290/go-gin-notes/app/repository"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repo = repository.NewRepo()

var (
	ErrUnmarshal = errors.New("unable to decode into struct")
	ErrFileRead  = errors.New("error reading the file")
)

func ConnectDatabase() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("config/")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration config.Configurations
	var err error

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(ErrFileRead, err)

	}

	// Set undefined variables
	//viper.SetDefault("database.dbname", "notes")
	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalln(ErrUnmarshal, err)

	}

	Repo.DbHost = configuration.Database.DBHost
	Repo.DbPass = configuration.Database.DBPassword
	Repo.DbUser = configuration.Database.DBUser
	Repo.DbPort = configuration.Database.DBPort
	Repo.DbName = configuration.Database.DBName

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
