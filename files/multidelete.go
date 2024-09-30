package files

import (
	"nat/auth"
	"net/http"
	"os"
	"strings"
	"sync"
)

func deleteFile(wg *sync.WaitGroup, fileName string, filesNotFound *[]string, sliceMutex *sync.Mutex) {
	defer wg.Done()
	filePath := "./" + fileName
	err := os.Remove(filePath)
	if err != nil {
		sliceMutex.Lock()
		*filesNotFound = append(*filesNotFound, fileName)
		sliceMutex.Unlock()
	}
}

func MultiDelete(w http.ResponseWriter, r *http.Request) {
	ah := r.Header.Get("Authorization")

	err := auth.ValidateAuthorizationHeader(ah)
	if err != nil {
		http.Error(w, "Could not validate token", http.StatusUnauthorized)
		return
	}

	query := r.URL.Query()
	filesList := query["file"]
	if len(filesList) == 0 {
		http.Error(w, "No files listed", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	filesNotFound := []string{}
	var sliceMutex sync.Mutex

	for _, fileName := range filesList {
		wg.Add(1)
		deleteFile(&wg, fileName, &filesNotFound, &sliceMutex)
	}

	wg.Wait()

	if len(filesNotFound) > 0 {
		w.Header().Set("Files-Not-Found", strings.Join(filesNotFound, ", "))
	}

	http.Error(w, "Files deleted succesfully", http.StatusOK)

}
