package auxfiles

// import the 1 modules we need
import "errors"

type ScanConfig struct {
}

type ScanType string

// Declare related constants for each security type
const (
	All                   ScanType = "All"
	CrossSiteScripting    ScanType = "CrossSiteScripting"
	SensitiveDataExposure ScanType = "SensitiveDataExposure"
	SQLInjection          ScanType = "SQLInjection"
)

func (ScanConfig) IsValid(sc string) error {
	scc := ScanType(sc)
	switch scc {
	case All, CrossSiteScripting, SensitiveDataExposure, SQLInjection:
		return nil
	}

	return errors.New("Invalid scan type")
}
