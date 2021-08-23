package dhlog

import "log"

func LogOutput(serviceName string, message string) {
	log.Printf("%v - %v", serviceName, message)
}

func NewLogger() Logger {
	log := LoggerAdapter(LogOutput)
	return log
}
