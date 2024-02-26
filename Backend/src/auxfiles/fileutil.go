package auxfiles

// import the 4 modules we need
import (
	"bufio"
	"mime/multipart"
)

type FileUtil struct {
}

func (FileUtil) ReadLines(fileHeader *multipart.FileHeader) (lines []string, err error) {
	// Open the file associated with the FileHeader
	inFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	// Create a new scanner to read from the file
	scanner := bufio.NewScanner(inFile)

	// Iterate over each line in the file
	for scanner.Scan() {
		// Append the line to the lines slice
		lines = append(lines, scanner.Text())
	}

	// Check if any errors occurred during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
