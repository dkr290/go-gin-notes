package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dkr290/go-devops/go-gin-notes/helpers"
	"github.com/dkr290/go-devops/go-gin-notes/models"
	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {

	notes := models.NotesAll()
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"notes/new.html",
		nil,
	)

}

func NotesCreate(c *gin.Context) {

	name := c.PostForm("name")
	content := c.PostForm("content")
	models.NoteCreate(name, content)

	c.Redirect(http.StatusMovedPermanently, "notes")

}

func NotesShow(c *gin.Context) {

	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	note := models.NotesFind(id)
	c.HTML(
		http.StatusOK,
		"notes/show.html",
		gin.H{
			"notes": note,
		},
	)
}

func NotesEditPage(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	note := models.NotesFind(id)
	c.HTML(
		http.StatusOK,
		"notes/edit.html",
		gin.H{
			"notes": note,
		},
	)
}

func NotesUpdate(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	note := models.NotesFind(id)
	name := c.PostForm("name")
	content := c.PostForm("content")
	note.NotesUpdate(name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes/"+idstr)
}

func NotesDelete(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	models.NotesMarkDelete(id)
	c.Redirect(http.StatusSeeOther, "/notes")

}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm-password")

	emailExists := models.UserCheckAvailability(email)
	if !emailExists {
		c.HTML(
			http.StatusIMUsed,
			"/home/signup.html",
			gin.H{
				"alert": "Email already exists",
			},
		)

		return
	}

	if password != confirmPassword {
		c.HTML(
			http.StatusIMUsed,
			"/home/signup.html",
			gin.H{
				"alert": "Passwords does not match",
			},
		)

		return
	}

	user := models.UserCreate(email, password)
	if user.ID == 0 {
		c.HTML(
			http.StatusIMUsed,
			"/home/signup.html",
			gin.H{
				"alert": "Unable to  create the user",
			},
		)

	} else {
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
