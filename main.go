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
		input = c.Param("code")
		if len(input) <= 0 {
			c.IndentedJSON(http.StatusBadRequest, "Enter a valid search param.")
			return
		}
	}

	convertedInput, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Enter a valid search param.")
		return
	}
	code := int16(convertedInput)

	result, err := getCode(code)


	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func searchByTitle(c *gin.Context) {
	input := c.Query("title")

	if len(input) <= 0 {
		input = c.Param("title")
		if len(input) <= 0 {
			c.IndentedJSON(http.StatusBadRequest, "Enter a search param.")
		}
	}

	result, err := findCodeByTitle(input)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()
	router.GET("/codes", codeList)

	router.GET("/codes/code/:code", searchByCode)
	router.GET("/codes/:code", searchByCode)

	router.GET("/codes/title", searchByTitle)
	router.GET("/codes/title/:title", searchByTitle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}