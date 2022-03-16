package logging

import (
	"fmt"
	"time"
)

var currentLevel Level

func SetLogLevel(level Level) {
	currentLevel = level
}

func printMessage(message string) {
	fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), message)
}

func Debug(message string) {
	if currentLevel < DEBUG {
		return
	}
	printMessage(message)
}

func Info(message string) {
	if currentLevel < INFO {
		return
	}
	printMessage(message)
}

func Warn(message string) {
	if currentLevel < WARN {
		return
	}
	printMessage(message)
}

func Error(message string) {
	if currentLevel < ERROR {
		return
	}
	printMessage(message)
}

func Fatal(message string) {
	if currentLevel < FATAL {
		return
	}
	printMessage(message)
}

func Trace(message string) {
	if currentLevel < TRACE {
		return
	}

	printMessage(message)
}
