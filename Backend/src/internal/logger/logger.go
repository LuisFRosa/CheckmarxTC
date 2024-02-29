package logger

// import the 2 modules we need
import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) LogInfo(message string) {
	l.infoLogger.Println(message)
}

func (l *Logger) LogWarning(message string) {
	l.warningLogger.Println(message)
}

func (l *Logger) LogError(message string) {
	l.errorLogger.Println(message)
}
