package files

import (
	"archive/zip"
	"fmt"
	"io"
	"nat/auth"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
)

func addFileToZip(wg *sync.WaitGroup, zipMutex *sync.Mutex, zipWriter *zip.Writer, fileName string, filesNotFound *[]string, sliceMutex *sync.Mutex) {
	defer wg.Done()
	f, err := os.Open(fileName)
	if err != nil {
		sliceMutex.Lock()
		*filesNotFound = append(*filesNotFound, fileName)
		sliceMutex.Unlock()
		return
	}

	defer f.Close()

	zipMutex.Lock()
	defer zipMutex.Unlock()

	wr, err := zipWriter.Create(fileName)
	if err != nil {
		return
	}

	_, err = io.Copy(wr, f)
	if err != nil {
		return
	}
}

func MultiDownload(w http.ResponseWriter, r *http.Request) {
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
	zipName := uuid.New().String() + ".zip"
	mdl, err := os.Create(zipName)
	if err != nil {
		http.Error(w, "Could not create zip file", http.StatusInternalServerError)
		return
	}
	defer os.Remove("./" + zipName)

	zipWriter := zip.NewWriter(mdl)

	filesNotFound := []string{}

	var wg sync.WaitGroup
	var zipMutex sync.Mutex
	var sliceMutex sync.Mutex

	for _, fileName := range filesList {
		wg.Add(1)
		go addFileToZip(&wg, &zipMutex, zipWriter, fileName, &filesNotFound, &sliceMutex)
	}

	wg.Wait()

	err = zipWriter.Close()
	if err != nil {
		http.Error(w, "Error finalizing zip", http.StatusInternalServerError)
		return
	}

	err = mdl.Close()
	if err != nil {
		http.Error(w, "Error closing zip", http.StatusInternalServerError)
		return
	}

	// send zip
	file, err := os.Open(zipName)
	if err != nil {
		http.Error(w, "Error opening generated zip", http.StatusInternalServerError)
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

	w.Header().Set("Content-Disposition", "attachment; filename="+zipName)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileStat.Size()))
	if len(filesNotFound) > 0 {
		w.Header().Set("Files-Not-Found", strings.Join(filesNotFound, ", "))
	}
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Could not send zip", http.StatusInternalServerError)
		return
	}
}
