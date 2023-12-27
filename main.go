package main

import (
	"fmt"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func main() {

	// Define handlers for file upload and download
	// http.HandleFunc("/upload", uploadFormHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/download/", downloadHandler)

	ip, err := getWLANIPAddress()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("WLAN IP address: %s\n", ip)
	// Start the server and listen on port 8080
	fmt.Println("File-sharing server listening on :8080")
	qrCode, err := qrcode.New(ip.To16().String()+":8080", qrcode.Medium)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	// Print QR code to the terminal
	fmt.Println(string(qrCode.ToString(false)))
	err = http.ListenAndServe(ip.To16().String()+":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
