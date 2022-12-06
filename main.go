package main

import (
	"log"
	"net/http"

	"github.com/dkr290/go-devops/go-gin-notes/controllers"
	"github.com/dkr290/go-devops/go-gin-notes/middlewares"
	"github.com/dkr290/go-devops/go-gin-notes/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	//loading static files
	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")
	models.ConnectDatabase()
	models.DbMigrate()

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))
	r.Use(middlewares.AuthenticateUser())

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)
	r.GET("/notes/:id", controllers.NotesShow)
	r.GET("/notes/edit/:id", controllers.NotesEditPage)
	r.POST("/notes/:id", controllers.NotesUpdate)
	r.DELETE("/notes/:id", controllers.NotesDelete)

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.POST("/logout", controllers.Logout)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":     "Notes Application",
			"logged_in": ctx.GetUint64("user_id") > 0,
		})
	})

	log.Println("Server is started!")
	r.Run("127.0.0.1:8080") // Default is 8080
}
