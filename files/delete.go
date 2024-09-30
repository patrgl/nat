package files

import (
	"nat/auth"
	"net/http"
	"os"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	ah := r.Header.Get("Authorization")

	err := auth.ValidateAuthorizationHeader(ah)
	if err != nil {
		http.Error(w, "Could not validate token", http.StatusUnauthorized)
		return
	}

	fileName := r.PathValue("filename")
	filePath := "./" + fileName

	err = os.Remove(filePath)
	if err != nil {
		http.Error(w, "Could not delete file", http.StatusBadRequest)
		return
	}

	http.Error(w, "File deleted succesfully", http.StatusOK)

}
