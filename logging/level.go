package logging

type Level uint32

const (
	OFF = iota
	TRACE
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	ALL
)
