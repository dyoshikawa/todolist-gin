package handler

import (
	"github.com/dyoshikawa/todolist-gin/infrastructure/mysql/mysql_util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	passwordConfirmation := c.PostForm("password_confirmation")
	if password != passwordConfirmation {
		c.Redirect(http.StatusSeeOther, "/sign_up")
		return
	}
	_, _ = mysql_util.Db.Exec("INSERT INTO users(email, password) VALUES(?, ?)", email, password)
	c.Redirect(http.StatusSeeOther, "/home")
}

func SignIn() {
	_, _ = mysql_util.Db.Exec("INSERT INTO users() VALUES(?, ?)")
}
