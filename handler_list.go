package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	// Open the directory
	dir, err := os.Open(fileDirectory)
	if err != nil {
		http.Error(w, "Error opening directory", http.StatusInternalServerError)
		return
	}
	defer dir.Close()

	// Read the contents of the directory
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		http.Error(w, "Error reading directory contents", http.StatusInternalServerError)
		return
	}

	// Prepare an HTML list of clickable links
	var fileListHTML strings.Builder
	fileListHTML.WriteString("<ul>")
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue // Skip directories
		}
		fileName := fileInfo.Name()
		fileLink := fmt.Sprintf(`<li><a href="/download/%s">%s</a></li>`, fileName, fileName)
		fileListHTML.WriteString(fileLink)
	}
	fileListHTML.WriteString("</ul>")

	// Serve the HTML list as a response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fileListHTML.String()))
}
