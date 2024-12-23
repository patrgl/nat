package main

import (
	"fmt"
	"nat/files"
	"net/http"
	"os"
)

func main() {

	os.Setenv("token", "thisisatoken")
	os.Setenv("baseurl", "localhost:8080")

	baseUrl := os.Getenv("baseurl")

	mux := http.NewServeMux()

	mux.HandleFunc("PUT /api/upload", files.Upload)
	mux.HandleFunc("GET /api/download/{filename}", files.Download)
	mux.HandleFunc("GET /api/multi-download", files.MultiDownload)
	mux.HandleFunc("DELETE /api/delete/{filename}", files.Delete)
	mux.HandleFunc("DELETE /api/multi-delete", files.MultiDelete)
	mux.HandleFunc("PUT /api/multi-upload", files.MultiUpload)

	fmt.Printf("Serving on %s", baseUrl)
	err := http.ListenAndServe(baseUrl, mux)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to start server")
		return
	}

}
