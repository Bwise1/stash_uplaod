package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func upload(w http.ResponseWriter, r *http.Request) {
	// Extract filename with extension from the request URL
	filenameWithExtension := filepath.Base(r.URL.Path)

	// Get the filename without extension
	filenameWithoutExtension := strings.TrimSuffix(filenameWithExtension, filepath.Ext(filenameWithExtension))

	// Create a temporary file with the original extension
	tempFile, err := os.CreateTemp("temp", filenameWithoutExtension+"-*"+filepath.Ext(filenameWithExtension))

	if err != nil {
		log.Println("Error creating temp file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(tempFile, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	wgetURL := "http://" + r.Host + "/download/" + filepath.Base(tempFile.Name())
	log.Println("File uploaded successfully ", wgetURL)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded successfully\nDownload with wget %s ", wgetURL)
	defer tempFile.Close()
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(r.URL.Path)
	log.Println("Downloading file ", fileName)

	file, err := os.Open("temp/" + fileName)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set appropriate headers for download
	w.Header().Set("Content-Type", "application/octet-stream")              // Set generic binary content type
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName) // Set attachment disposition with filename

	// Stream the file content to the response
	io.Copy(w, file)
}
func main() {

	mux := http.NewServeMux()

	log.Println("Stash Upload")

	mux.HandleFunc("PUT /", upload)
	mux.HandleFunc("GET /download/", downloadHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
