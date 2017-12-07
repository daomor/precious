package logger

import (
	"os"
	"log"
)


type LogService struct {}

const INFO = "info"
type Test log.Logger

var logdata *LogData

func init() {
	logdata = &LogData{}
}

func NewLogger() *LogService {
	return &LogService{}
}

/*
	TODO:
	- Document the syntax.
	- Tidy files.
	- Add to repo
	-
 */

/**
 * LOG CREATION
 */

func (*LogService) CreateCustom(logType string, filepath string) {
	createLog(logType, filepath)
}

func (*LogService) Create(filepath string) {
	createLog(INFO, filepath)
}

func createLog(logType string, filepath string) {

	// Create file.
	file, err := os.OpenFile(filepath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	// Check for error when making the file.
	if(err != nil) {
		log.Fatal("Error opening file: " + filepath)
	}

	lg := &MyLog{
		log: file,
		Name: logType,
	}

	// Add the log to the array.
	logdata.Logs = append(logdata.Logs, lg)
}

func (*LogService) Log(logType string) *MyLog {
	return logdata.LogItem(logType)
}

func (*LogService) Close() {
	for _, v := range logdata.Logs {
		v.log.Close()
	}
}


/*
	Syntax
	------

	// Creates a log file at the path.
	logger.Create(logger.Info, path)
	logger.Create('testlog', path)

	// Log to access log.
	logger.Message()

	// Log to custom file.
	logger.Custom('testlog', 'message')
 */