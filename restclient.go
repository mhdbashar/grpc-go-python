package main

import (
	"fmt"
	"net/http"
    "encoding/json"
	"io/ioutil"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
    Body   string `json:"body"`
}

func main() {
    // Create an HTTP client
    client := &http.Client{}

    // Define the API endpoint URL
    apiUrl := "http://10.10.10.80:5586/11"


    // Send an HTTP GET request
    resp, err := client.Get(apiUrl)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    body,err := ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println("Error:", err)
	return
	 }
    var post Post
    err = json.Unmarshal(body, &post)
    if err != nil {
        fmt.Println("Error:", err)
	return
         }

    fmt.Println("JSON data:", post)

    // defer resp.Body.Close()

    // Check the HTTP status code
    if resp.StatusCode != http.StatusOK {
        fmt.Println("HTTP request failed with status:", resp.Status)
        return
    }

    // Decode the JSON response
    // var post Post
    // err = json.NewDecoder(resp.Body).Decode(&post)
    // if err != nil {
    //     fmt.Println("JSON decoding error:", err)
    //     return
    // }

    // Print the parsed data
    fmt.Printf("Post ID: %d\n", post.ID)
    fmt.Printf("Title: %s\n", post.Title)
    fmt.Printf("Body: %s\n", post.Body)
}

