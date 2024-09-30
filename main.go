package main

import (
	"fmt"
	"nat/files"
	"net/http"
	"os"
)

func main() {

	os.Setenv("token", "thisisatoken")

	mux := http.NewServeMux()

	mux.HandleFunc("PUT /api/upload", files.Upload)
	mux.HandleFunc("GET /api/download/{filename}", files.Download)
	mux.HandleFunc("GET /api/multi-download", files.MultiDownload)
	mux.HandleFunc("DELETE /api/delete/{filename}", files.Delete)
	mux.HandleFunc("DELETE /api/multi-delete", files.MultiDelete)

	fmt.Println("Serving on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err)
	}

}
