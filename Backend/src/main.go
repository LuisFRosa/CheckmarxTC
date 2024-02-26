package main

// import the 6 modules we need
import (
	auxFiles "Backend/src/auxfiles"

	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
)

type ScanResult struct {
	Errors []string `json:"errors"`
}

func main() {
	http.HandleFunc("/scan", scanHandler)
	http.ListenAndServe(":8000", nil)
}

func scanHandler(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Encode ScanResult to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20) // maxMemory 10 MB

	files := r.MultipartForm.File["files"]
	scanConfig := r.URL.Query().Get("scanConfig")

	// Validate the scan configuration
	if ((auxFiles.ScanConfig).IsValid(auxFiles.ScanConfig{}, scanConfig) != nil) {
		http.Error(w, "Invalid scan configuration", http.StatusBadRequest)
		return
	}

	fmt.Println("Scan configuration: " + scanConfig)
	fmt.Println("Number of files: " + strconv.Itoa((len(files))))

	var errors []string
	for _, fileHeader := range files {
		// Get the filename
		filename := fileHeader.Filename
		fmt.Println("Filename:", filename)

		// Check for specific types of errors based on scanConfig value
		errors = append(scanfile(scanConfig, fileHeader, filename), errors...)
		if errors == nil {
			http.Error(w, "Failed to scan file", http.StatusInternalServerError)
			return
		}
	}

	result := ScanResult{Errors: errors}
	json.NewEncoder(w).Encode(result)
}

func scanfile(scan string, fileHeader *multipart.FileHeader, fileName string) []string {

	var errorsOfFile = []string{}
	lines, err := (auxFiles.FileUtil).ReadLines(auxFiles.FileUtil{}, fileHeader)
	if err != nil {
		return nil
	}

	for index, line := range lines {
		errorsOfLine := (auxFiles.SecurityUtil).CheckLine(auxFiles.SecurityUtil{}, scan, fileName, line, index)
		for i := 0; i < len(errorsOfLine); i++ {
			errorsOfFile = append(errorsOfFile, errorsOfLine[i])
		}
	}

	return errorsOfFile
}
