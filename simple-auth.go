package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jairoandre/simple-auth/model"
)

func postUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user = addUser(user)
	c.JSON(200, user)
}

func getUsers(c *gin.Context) {
	users := findAllUsers()
	c.JSON(200, users)
}

func server() {
	r := gin.Default()
	r.POST("/user", postUser)
	r.GET("/user", getUsers)
	r.GET("/ping", func (c *gin.Context){
		c.JSON(200, gin.H{"message": "pong",})
	})
	r.Run()
}

func main() {
	server()
}
