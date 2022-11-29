package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func codeList(c *gin.Context) {
	list := getCodes()
	c.IndentedJSON(http.StatusOK, list) 
}

func main() {
	fmt.Println(getTitle(400))

	router := gin.Default()
	router.GET("/codes", codeList)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}