package logger 

import (
	"log"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO 
	WARM
	ERROR
)

var currentLevel = INFO 

func SetLogLevel(level LogLevel) {
	currentLevel = level
}

func logMessage(level LogLevel, msg string) {
	if level >= currentLevel {
		timestamp := time.Now().Format('2006-01-02 15:04:05')
		log.Printf("[%s] %s: %s\n", timestamp, levelToString(level), msg)

	}
}

func levelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WANR:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	
	}

}

func Debug(msg string) {
	logMessage(DEBUG, msg)
}

func Info(msg string){
	logMessage(INFO, msg)
}

func WARN(msg string)  {
	logMessage(WARN, msg)
}

func Error(msg string) {
	logMessage(ERROR, msg)
}

func main()  {
	//Example of usage
	SetLogLevel(DEBUG)
	Info("This is na info message")
	Debug("This is a debug message")
	Error("This is an error message")
}