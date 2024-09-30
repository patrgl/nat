package files

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"nat/auth"
)

func Download(w http.ResponseWriter, r *http.Request) {
	ah := r.Header.Get("Authorization")

	err := auth.ValidateAuthorizationHeader(ah)
	if err != nil {
		http.Error(w, "Could not validate token", http.StatusUnauthorized)
		return
	}

	fileName := r.PathValue("filename")
	filePath := "./" + fileName

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Could not find file", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileHeader := make([]byte, 512)
	file.Read(fileHeader)
	contentType := http.DetectContentType(fileHeader)
	fileStat, err := file.Stat()
	if err != nil {
		http.Error(w, "Error reading file size", http.StatusInternalServerError)
		return
	}

	file.Seek(0, 0)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileStat.Size()))
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Could not send file", http.StatusInternalServerError)
		return
	}
}
