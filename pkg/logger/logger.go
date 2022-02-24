//
package logger

import (
	"log"
	"os"
)

type Loggers struct {
	LogVerbose *log.Logger
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
}

// Instantiates new log instances for logging different levels of information
func CreateLoggers(log_file string) *Loggers {
	file, err := os.OpenFile(log_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	loggers := Loggers{
		log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)}

	return &loggers
}
