package views

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sayhello_view(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}

func FormGet_view(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{})
}

func FormPost_view(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}
