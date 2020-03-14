package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/users", getUsers)

	r.Run(":" + os.Getenv("PORT"))
}

func getUsers(c *gin.Context) {
	id := c.GetInt("id")
	user := struct {
		Id 		int 	`json:"id"` // аннотация используется для указания как обозначать поля в json
		Name 	string 	`json:"name"`
	}{id, "Сергей"}

	c.JSON(200, user)
}
