package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const downloadDirectory = "./"

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the filename from the URL path

	fileName := r.URL.Path[len("/download/"):]
	// Open the requested file
	filePath := downloadDirectory + fileName
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the Content-Disposition header to trigger download
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")

	// Copy the file content to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error sending file", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 200)
}
