package logger

// import the 2 modules we need
import (
	"log"
	"os"
)

// LogLevel represents the logging level
type LogLevel int

const (
	// LogLevelInfo represents the info logging level
	LogLevelInfo LogLevel = iota
	// LogLevelWarning represents the warning logging level
	LogLevelWarning
	// LogLevelError represents the error logging level
	LogLevelError
)

// Logger represents the logger
type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

// NewLogger creates a new instance of Logger
func NewLogger() *Logger {
	return &Logger{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// LogInfo logs a message with info level
func (l *Logger) LogInfo(message string) {
	l.infoLogger.Println(message)
}

// LogWarning logs a message with warning level
func (l *Logger) LogWarning(message string) {
	l.warningLogger.Println(message)
}

// LogError logs a message with error level
func (l *Logger) LogError(message string) {
	l.errorLogger.Println(message)
}
