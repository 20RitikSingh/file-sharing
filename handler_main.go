package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

const uploadFormHTML = `
<!DOCTYPE html>
<html>
<head>
	<title>ShareGo</title>
</head>
<body>
	<h1>ShareGo</h1>
	<h2>Upload file:</h2>
	<form enctype="multipart/form-data" action="/upload" method="post">
		<input type="file" name="uploadfile" />
		<input type="submit" value="Upload" />
	</form>
	<br/>
	<h2>Files in current directory:</h2>
	{{.FileList}}
</body>
</html>
`

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data, including file uploads
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Get the file from the form data
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create a new file on the server to store the uploaded file
		f, err := os.Create(handler.Filename)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Copy the file to the new file on the server
		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Open the directory
	dir, err := os.Open(".")
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
	var fileListHTML string
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue // Skip directories
		}
		fileName := fileInfo.Name()
		fileListHTML += fmt.Sprintf(`<li><a href="/download/%s">%s</a></li>`, fileName, fileName)
	}

	// Render the HTML form with the file list
	tmpl, err := template.New("uploadForm").Parse(uploadFormHTML)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"FileList": fileListHTML,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
