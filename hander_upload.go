package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// const uploadFormHTML = `
// <!DOCTYPE html>
// <html>
// <head>
// 	<title>File Upload</title>
// </head>
// <body>
// 	<form enctype="multipart/form-data" action="/recieve" method="post">
// 		<input type="file" name="uploadfile" />
// 		<input type="submit" value="Upload" />
// 	</form>
// </body>
// </html>
// `

// func uploadFormHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.New("uploadForm").Parse(uploadFormHTML)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }

func uploadHandler(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "<p>File uploaded successfully: %s</p>", handler.Filename)
	fmt.Fprintf(w, "<a href=\"/\">Back</a>")

}
