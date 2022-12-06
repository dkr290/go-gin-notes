package models

import (
	"log"
	"time"

	"github.com/dkr290/go-devops/go-gin-notes/helpers"
)

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	UserName  string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserCheckAvailability(email string) bool {
	var user User
	Repo.DB.Find(&user, email)
	return (user.ID == 0) // id is 0 the email has not been used and available
}

func UserCreate(email, password string) *User {
	hashPassord, err := helpers.HashPassword(password)
	if err != nil {
		log.Fatal("The password cannot be hashed")
	}
	entry := User{
		UserName: email,
		Password: hashPassord,
	}
	Repo.DB.Create(&entry)
	return &entry
}

func UserFind(sessionID uint64) *User {

	var user User
	Repo.DB.Find(&user, sessionID)
	return &user

}

func UserCheck(email, password string) *User {

	var user User
	Repo.DB.Find(&user, "username")
	if user.ID == 0 {
		return nil
	}

	match := helpers.CheckPasswordHash(password, user.Password)
	if match {
		return &user
	} else {
		return nil
	}
}
