package controllers

import (
	"net/http"

	"github.com/dkr290/go-devops/go-gin-notes/helpers"
	"github.com/dkr290/go-devops/go-gin-notes/models"
	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"home/login.html",
		nil,
	)
}

func SignupPage(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"home/signup.html",
		nil,
	)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := models.UserCheck(email, password)
	if user != nil {
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(
			http.StatusOK,
			"home/login.html",
			gin.H{
				"alert": "Email and/or password mismatch",
			},
		)
	}
}

func Logout(c *gin.Context) {
	//clear the session
	helpers.SessionClear(c)
	c.HTML(
		http.StatusOK,
		"home/logout.html",
		gin.H{
			"alert": "Logged Out",
		},
	)

}
