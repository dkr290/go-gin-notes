package middlewares

import (
	"github.com/dkr290/go-devops/go-gin-notes/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("id")
		var user *models.User
		userPresent := true

		if sessionID == nil {
			userPresent = false
		} else {
			user = models.UserFind(sessionID.(uint64))
			userPresent = (user.ID > 0)
		}

		if userPresent {
			c.Set("user_id", user.ID)
			c.Set("email", user.UserName)
		}
		c.Next()

	}
}
