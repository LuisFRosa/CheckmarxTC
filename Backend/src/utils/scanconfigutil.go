package utils

// import the 1 modules we need
import "errors"

type ScanType string

// Declare related constants for each security type
const (
	All                   ScanType = "All"
	CrossSiteScripting    ScanType = "CrossSiteScripting"
	SensitiveDataExposure ScanType = "SensitiveDataExposure"
	SQLInjection          ScanType = "SQLInjection"
)

func IsValid(sc string) error {
	scc := ScanType(sc)
	switch scc {
	case All, CrossSiteScripting, SensitiveDataExposure, SQLInjection:
		return nil
	}

	return errors.New("Invalid scan type")
}
