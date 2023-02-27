package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type Bookmark struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Url     string `json:"url"`
}

var fakeToken = tokenResponse{
	AccessToken:  "you-are-good",
	Expires:      1924905600000,
	RefreshToken: "still-good",
	Scope:        "all",
	TokenType:    "very-special",
}

func returnToken(c *gin.Context) {
	c.JSON(200, fakeToken)
}

func saveBookmark(c *gin.Context) {
	var bookmark Bookmark
	err := c.BindJSON(&bookmark)
	if err != nil {
		fmt.Printf("invalid payload")
		log.Fatal(err)
	}

	c.Status(200)

	json_data, err := json.Marshal(bookmark)

	if err != nil {
		log.Fatal(err)
	}

	targetUrl := os.Getenv("TARGET_URL")
	// e.g. to save to a pocketbase collection add TARGET_URL=https://{pocketbase-instance-address}/api/collections/{collection-name}/records

	resp, err := http.Post(targetUrl, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}

func main() {
	router := gin.Default()
	router.POST("/oauth/v2/token", returnToken)
	router.POST("/api/entries.json", saveBookmark)

	router.Run()
}
