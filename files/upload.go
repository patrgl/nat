package files

import (
	"io"
	"net/http"
	"os"

	"nat/auth"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	ah := r.Header.Get("Authorization")

	err := auth.ValidateAuthorizationHeader(ah)
	if err != nil {
		http.Error(w, "Could not validate token", http.StatusUnauthorized)
		return
	}

	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Could not get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, "Could not create file", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Could not save file", http.StatusBadRequest)
		return
	}

	http.Error(w, "Success", http.StatusAccepted)
}
