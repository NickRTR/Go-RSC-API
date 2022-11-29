package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func codeList(c *gin.Context) {
	list := getCodes()
	c.IndentedJSON(http.StatusOK, list) 
}

func searchByCode(c *gin.Context) {
	input := c.Query("code")

	if len(input) <= 0 {
		c.IndentedJSON(http.StatusBadRequest, "Enter a search param.")
	}

	convertedIntput, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Enter a valid search param.")
		return
	}
	code := int16(convertedIntput)

	title, err := getTitle(code)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, title)
}

func main() {
	router := gin.Default()
	router.GET("/codes", codeList)
	router.GET("/code", searchByCode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}