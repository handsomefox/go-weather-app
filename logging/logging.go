package logging

import (
	"fmt"
	"time"
)

type LogLevel uint8

const (
	_ = iota
	OFF
	TRACE
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	ALL
)

type Logger struct {
	Level LogLevel
}

func (logger Logger) Debug(message string) {
	if logger.Level < DEBUG {
		return
	}
	printMessage(message, "DEBUG")
}

func (logger Logger) Info(message string) {
	if logger.Level < INFO {
		return
	}
	printMessage(message, "INFO")
}

func (logger Logger) Warn(message string) {
	if logger.Level < WARN {
		return
	}
	printMessage(message, "WARN")
}

func (logger Logger) Error(message string) {
	if logger.Level < ERROR {
		return
	}
	printMessage(message, "ERROR")
}

func (logger Logger) Fatal(message string) {
	if logger.Level < FATAL {
		return
	}
	printMessage(message, "FATAL")
}

func (logger Logger) Trace(message string) {
	if logger.Level < TRACE {
		return
	}

	printMessage(message, "TRACE")
}

func printMessage(message string, level string) {
	fmt.Printf("[%s %s] %s\n", time.Now().Format(time.RFC3339), level, message)
}
