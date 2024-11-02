package logger

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitLogger(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name        string
		logLevel    string
		logPath     *string
		logFormat   LogFormat
		maxSize     int
		maxBackups  int
		maxAge      int
		expectPanic bool
	}{
		{"InfoLevelWithConsoleFormat", "info", nil, ConsoleFormat, 10, 3, 7, false},
		{"DebugLevelWithJSONFormat", "debug", nil, JSONFormat, 5, 1, 2, false},
		{"PanicOnInvalidFormat", "info", nil, LogFormat("invalid"), 5, 1, 2, true},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// If the test case is supposed to panic, assert that it does
			if tc.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			// Initialize the logger with the test case parameters
			InitLogger(tc.logLevel, tc.logPath, tc.logFormat, tc.maxSize, tc.maxBackups, tc.maxAge)

			// Assert that the logger is not nil
			if GetLogger() == nil {
				t.Errorf("Logger instance is nil")
			}

			// Add more assertions as needed to validate the behavior
		})
	}
}

func TestGetLogger(t *testing.T) {
	// Initialize the logger once before the test
	InitLogger("info", nil, ConsoleFormat, 10, 3, 7)

	// Test the GetLogger function
	logger := GetLogger()
	if logger == nil {
		t.Errorf("GetLogger returned nil")
	}

	// Assert that the logger level is set correctly
	if logger.GetLevel() != logrus.InfoLevel {
		t.Errorf("Logger level is not set to info, got: %v", logger.GetLevel())
	}

	// Add more assertions as needed to validate the behavior
}

// TestLogFileWrite 测试日志文件的写入
func TestLogFileWrite(t *testing.T) {
	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "test-log-")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Define the log file path
	logFilePath := filepath.Join(tempDir, "test.log")

	// Initialize the logger with the log file path
	InitLogger("info", &logFilePath, ConsoleFormat, 10, 3, 7)

	// Get the logger instance
	logger := GetLogger()
	if logger == nil {
		t.Errorf("GetLogger returned nil")
	}

	// Log a message
	testMessage := "This is a test log message"
	logger.Info(testMessage)

	// Read the log file content
	content, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		t.Errorf("Failed to read log file: %v", err)
	}

	// Check if the log message is in the file content
	if !strings.Contains(string(content), testMessage) {
		t.Errorf("Log message not found in the log file: %s", testMessage)
	}
}

// To run the tests, use the command `go test` in the terminal within the package directory.
