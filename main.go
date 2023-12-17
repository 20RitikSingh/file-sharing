package main

import (
	"fmt"
	"net/http"
	"os"
)

const uploadDirectory = "./uploads/"
const fileDirectory = "./"

func main() {
	// Create the uploads directory if it doesn't exist
	err := os.MkdirAll(uploadDirectory, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating upload directory:", err)
		return
	}

	// Define handlers for file upload and download
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/download/", downloadHandler)

	// Start the server and listen on port 8080
	fmt.Println("File-sharing server listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
