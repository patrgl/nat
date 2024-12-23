package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"sync"

	"nat/auth"
)

func WriteFile(wg *sync.WaitGroup, filesNotWriten *[]string, filesNotWritenMutex *sync.Mutex, handler *multipart.FileHeader) {
	defer wg.Done()

	file, err := handler.Open()
	if err != nil {
		filesNotWritenMutex.Lock()
		*filesNotWriten = append(*filesNotWriten, handler.Filename)
		filesNotWritenMutex.Unlock()
		return
	}
	defer file.Close()

	fileNameToUse := handler.Filename
	_, err = os.Stat(handler.Filename)
	if err == nil {
		fileNameToUse = fmt.Sprintf("%s (1)", handler.Filename)
	}

	dst, err := os.Create(fileNameToUse)
	if err != nil {
		filesNotWritenMutex.Lock()
		*filesNotWriten = append(*filesNotWriten, handler.Filename)
		filesNotWritenMutex.Unlock()
		return
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		filesNotWritenMutex.Lock()
		*filesNotWriten = append(*filesNotWriten, handler.Filename)
		filesNotWritenMutex.Unlock()
		return
	}

	return

}

func MultiUpload(w http.ResponseWriter, r *http.Request) {
	ah := r.Header.Get("Authorization")

	err := auth.ValidateAuthorizationHeader(ah)
	if err != nil {
		http.Error(w, "Could not validate token", http.StatusUnauthorized)
		return
	}

	r.ParseMultipartForm(32 << 20)
	files := r.MultipartForm.File["file"]
	if len(files) == 0 {
		http.Error(w, "No files in request body", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	var filesNotWrittenMutex sync.Mutex
	filesNotWritten := []string{}

	for _, handler := range files {
		wg.Add(1)
		go WriteFile(&wg, &filesNotWritten, &filesNotWrittenMutex, handler)
	}

	wg.Wait()

	if len(filesNotWritten) > 0 {
		w.Header().Set("Files-Not-Uploaded", strings.Join(filesNotWritten, ", "))
	}

	w.WriteHeader(http.StatusOK)
	return
}
