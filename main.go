package main

import (
	//"cli/cli"
	"fmt"
	"cli/cli"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	type album struct {
		ID     string  `json:"id"`
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float64 `json:"price"`
	}

	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

    c.IndentedJSON(http.StatusOK, albums)
}


func server(){
	    // Create a new Gin router
		router := gin.Default()

		// Define a route for handling POST requests at /guess
		router.POST("/guess", func(c *gin.Context) {
			// Read the request body
			body, err := c.GetRawData()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
				return
			}
	
			// Print the received data
			println("Received POST request data:", string(body))
		
			client_guess_string, err := jsonparser.GetString(body, "guess")
			if err != nil {
				fmt.Println("error: ", err)
				return
			}
			
			fmt.Println(guessHandler.GuessHandler(client_guess_string))
			c.JSON(http.StatusOK, gin.H{"message": "POST request received successfully"})
		})
	
		// Start the server on port 8080
		if err := router.Run("127.0.0.1:8080"); err != nil {
			println("Server error:", err)
		}
}


func main() {
	server()
	fmt.Println()
}