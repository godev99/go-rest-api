package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Name struct {
	Type       string `json:"Type"`
	Kind       string `json:"Kind"`
	Expression string `json:"Expression"`
}

type Names []Name

var allNames = Names{
	{
		Type:       "KeyVault",
		Kind:       "PublicPaas",
		Expression: "regexatrouver",
	},
	{
		Type:       "WebApp",
		Kind:       "PrivatePaas",
		Expression: "regexatrouver",
	},
	{
		Type:       "VirtualMachine",
		Kind:       "Iaas",
		Expression: "regexatrouver",
	},
}

func homeLink(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, "Welcome home!")
}

// createName adds a new name from JSON received in the request body.
func createName(c *gin.Context) {
	var newName Name

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newName); err != nil {
		return
	}

	// Add the new album to the slice.
	allNames = append(allNames, newName)
	c.IndentedJSON(http.StatusCreated, newName)
}

func getNameByType(c *gin.Context) {
	// Get the Type from the url
	myType := c.Param("Type")

	// Loop through the list of names, looking for
	// an name whose Type value matches the parameter.
	for _, a := range allNames {
		if a.Type == myType {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Type not found"})
}

// getNames responds with the list of all names
func getNames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allNames)
}

func main() {
	router := gin.Default()
	router.GET("/", homeLink)
	router.GET("/names", getNames)
	router.GET("/name/:Type", getNameByType)
	router.POST("/name", createName)
	router.Run()
}
