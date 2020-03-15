package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.POST("/user", postUser)

	r.Run(":" + os.Getenv("PORT"))
}

func getUsers(c *gin.Context) {
	paramId := c.Query("id")
	id, _ := strconv.Atoi(paramId)
	user := struct {
		Id 		int 	`json:"id"` // аннотация используется для указания как обозначать поля в json
		Name 	string 	`json:"name"`
	}{id, "Сергей"}

	c.JSON(200, user)
}

func postUser(c *gin.Context) {
	user := struct {
		Id 		int 	`json:"id"` // аннотация используется для указания как обозначать поля в json
		Name 	string 	`json:"name"`
	}{}
	c.BindJSON(&user)

	c.JSON(200, user)
}
