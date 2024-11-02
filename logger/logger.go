package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

// Logger is a global logger instance
var Logger *log.Logger

// LogFormat defines the types of log formats
type LogFormat string

const (
	JSONFormat    LogFormat = "json"
	TextFormat    LogFormat = "text"
	ConsoleFormat LogFormat = "console"
)

// getGoroutineID returns the ID of the current goroutine
func getGoroutineID() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := fmt.Sprintf("%s", buf[:n])
	id := idField[len("goroutine "):]
	space := strings.Index(id, " ")
	if space > 0 {
		return id[:space]
	}
	return ""
}

// CustomFormatter is a custom log formatter
type CustomFormatter struct{}

// Format formats the log entry
func (f *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
	prefix := fmt.Sprintf("[%s] [%s] [%s:%d] ", getGoroutineID(), entry.Time.Format("2006-01-02 15:04:05"), entry.Caller.File, entry.Caller.Line)
	formattedMessage := fmt.Sprintf("%s %s\n", prefix, entry.Message)
	return []byte(formattedMessage), nil
}

// InitLogger initializes the logger
//
// Parameters:
// - logLevel: The log level as a string (e.g., "info", "debug")
// - logPath: A pointer to a string representing the log file path
// - logFormat: The format of the log (JSON, Text, Console)
// - maxSize: The maximum size in megabytes for each log file
// - maxBackups: The maximum number of backup log files to retain
// - maxAge: The maximum number of days to retain old log files
func InitLogger(logLevel string, logPath *string, logFormat LogFormat, maxSize int, maxBackups int, maxAge int) {
	// Create a new logrus instance
	Logger = log.New()

	// Set the log level
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	Logger.SetLevel(level)

	// Set the default log format to console
	if logFormat == "" {
		logFormat = ConsoleFormat
	}

	if logFormat != ConsoleFormat && logFormat != JSONFormat && logFormat != TextFormat {
		panic("Invalid log format")
	}

	// Set the log output format
	switch logFormat {
	case JSONFormat:
		Logger.SetFormatter(&log.JSONFormatter{})
	case TextFormat:
		Logger.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	case ConsoleFormat:
		Logger.SetFormatter(&CustomFormatter{}) // Use custom format
	default:
		Logger.SetFormatter(&CustomFormatter{}) // Default fallback to custom format
	}

	// Enable caller tracing
	Logger.SetReportCaller(true)

	// Set the log output target
	if logPath == nil || *logPath == "" {
		Logger.SetOutput(os.Stdout)
	} else {
		// Configure lumberjack
		lj := &lumberjack.Logger{
			Filename:   filepath.Clean(*logPath),
			MaxSize:    maxSize,    // Maximum size in MB for each log file
			MaxBackups: maxBackups, // Maximum number of backup log files to retain
			MaxAge:     maxAge,     // Maximum number of days to retain old log files
			Compress:   true,       // Whether to compress old log files
		}

		Logger.SetOutput(lj)
	}
}

// GetLogger returns the global logger instance
func GetLogger() *log.Logger {
	return Logger
}
