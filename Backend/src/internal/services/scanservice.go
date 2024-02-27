package services

// import the 6 modules we need
import (
	"Backend/src/internal/logger"
	"Backend/src/utils"

	"errors"
	"mime/multipart"
	"strconv"
)

type ScanService struct {
}

// ScanHandler is a method of ScanService that handles scanning files with a given scan configuration
func (s *ScanService) ScanHandler(files []*multipart.FileHeader, scanConfig string) ([]string, error) {

	// Init logger instance
	log := logger.NewLogger()

	// Validate the scan configuration
	if utils.IsValid(scanConfig) != nil {
		log.LogError("Invalid scan configuration")
		return nil, errors.New("Invalid scan configuration")
	}

	log.LogInfo("Scan configuration > " + scanConfig)
	log.LogInfo("Number of files > " + strconv.Itoa((len(files))))

	var errorsToShow []string
	for _, fileHeader := range files {
		// Get the filename
		filename := fileHeader.Filename
		log.LogInfo("Filename > " + filename)

		// Check for specific types of errors based on scanConfig value
		errorsToShow = append(scanfile(scanConfig, fileHeader, filename), errorsToShow...)
		if errorsToShow == nil {
			log.LogError("Failed to scan file")
			return nil, errors.New("Failed to scan file")
		}
	}

	return errorsToShow, nil
}

func scanfile(scan string, fileHeader *multipart.FileHeader, fileName string) []string {

	var errorsOfFile = []string{}
	lines, err := utils.ReadLines(fileHeader)
	if err != nil {
		return nil
	}

	for index, line := range lines {
		errorsOfLine := utils.CheckLine(scan, fileName, line, index)
		for i := 0; i < len(errorsOfLine); i++ {
			errorsOfFile = append(errorsOfFile, errorsOfLine[i])
		}
	}

	return errorsOfFile
}
