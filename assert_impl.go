package assert

import "fmt"

type logFacade struct {
	t      TestingT
	logger errorLogger
}

func setupImpl(t TestingT, logger errorLogger) *Provider {
	return &Provider{&logFacade{t, logger}}
}

func logIfFalse(logFacade *logFacade, condition bool, format string, args ...interface{}) {
	if logFacade != nil && !condition {
		location := provideLocation(3)
		message := fmt.Sprintf(format, args...)
		logFacade.Log(location, message)
	}
}

func (logFacade *logFacade) Log(location *location, message string) {
	logFacade.t.Fail()
	logFacade.logger.Log(location, message)
}
