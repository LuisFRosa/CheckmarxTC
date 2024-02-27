package main

// import the 4 modules we need
import (
	"Backend/src/internal/logger"
	"Backend/src/internal/services"

	"encoding/json"
	"net/http"
)

type ScanResult struct {
	Errors []string `json:"errors"`
}

func scanner(w http.ResponseWriter, r *http.Request) {

	// Init logger instance
	log := logger.NewLogger()

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Encode ScanResult to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")

	// Only accept method POST
	if r.Method != http.MethodPost {
		log.LogError("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Max memory 10 MB
	r.ParseMultipartForm(10 << 20)

	files := r.MultipartForm.File["files"]
	scanConfig := r.URL.Query().Get("scanConfig")

	// Create an instance of ScanService
	scanService := services.ScanService{}
	listOfErrors, err := scanService.ScanHandler(files, scanConfig)
	if err != nil {
		log.LogError(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := ScanResult{Errors: listOfErrors}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/scan", scanner)
	http.ListenAndServe(":8080", nil)
}
