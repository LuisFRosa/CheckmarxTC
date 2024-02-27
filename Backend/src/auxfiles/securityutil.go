package auxfiles

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type SecurityUtil struct {
}

func (SecurityUtil) CheckLine(scanconfig string, fileName string, line string, index int) (errors []string) {

	// overwrite number of line
	nLine := strconv.Itoa(index + 1)

	// cross site scripting
	if (scanconfig == "All" || scanconfig == "CrossSiteScripting") &&
		(filepath.Ext(fileName) == ".html" || filepath.Ext(fileName) == ".js") &&
		strings.Contains(line, "Alert()") {

		errors = append(errors, "[Cross site scripting] in file \""+fileName+"\" on line "+nLine)
	}

	re := regexp.MustCompile(`"[^"]+"`)
	quotes := re.FindAllString(line, -1)
	for _, quote := range quotes {

		// sensitive data exposure.
		if (scanconfig == "All" || scanconfig == "SensitiveDataExposure") &&
			(strings.Contains(quote, "Checkmarx") ||
				strings.Contains(quote, "Hellman & Friedman") ||
				strings.Contains(quote, "$1.15b")) {

			errors = append(errors, "[Sensitive data exposure.] in file \""+fileName+"\" on line "+nLine)
		}

		// SQL injection
		if (scanconfig == "All" || scanconfig == "SQLInjection") &&
			strings.Contains(quote, "SELECT") &&
			strings.Contains(quote, "WHERE") &&
			strings.Contains(quote, "%s") {

			errors = append(errors, "[SQL injection] in file \""+fileName+"\" on line "+nLine)
		}
	}

	return errors
}
