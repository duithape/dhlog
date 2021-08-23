package dhlog

type LoggerAdapter func(serviceName string, message string)

func (lg LoggerAdapter) Log(serviceName string, message string) {
	lg(serviceName, message)
}
