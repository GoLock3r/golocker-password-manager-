package logger

import (
	"bufio"
	"os"
	"regexp"
	"testing"
)

func removeFiles() {
	os.Remove("testlogins.txt")
	os.Remove("testlogs.txt")
}

func TestCreateLoggers(t *testing.T) {
	logfile := "testlogs.txt"
	loggers := CreateLoggers(logfile)

	loggers.LogVerbose.Print("")

	_, err := os.Stat(logfile)
	if err != nil {
		t.Error("Expected existance of", logfile, "but no file was found.")
	}

	removeFiles()
}

func TestLoggerLevels(t *testing.T) {
	logfile := "testlogs.txt"
	loggers := CreateLoggers(logfile)

	loggers.LogVerbose.Print("")
	loggers.LogInfo.Print("")
	loggers.LogWarning.Print("")
	loggers.LogError.Print("")

	file, _ := os.Open(logfile)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("DEBUG|INFO|WARNING|ERROR")
	matches := 0
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			matches++
		}
	}

	if matches != 4 {
		t.Error("Expected 4 levels of data logging, returned", matches)
	}

	removeFiles()
}
