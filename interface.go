package dhlog

type Logger interface {
	Log(serviceName string, message string)
}
