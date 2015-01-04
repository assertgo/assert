package assert

import "fmt"

type Provider struct {
	logFacade *logFacade
}

type logFacade struct {
	t      TestingT
	logger errorLogger
}

func setupImpl(t TestingT, logger errorLogger) *Provider {
	return &Provider{&logFacade{t, logger}}
}

func (p *Provider) That(actual interface{}) *AnyType {
	return &AnyType{p.logFacade, actual}
}

func (p *Provider) ThatInt(actual int) *Int {
	return &Int{p.logFacade, actual}
}

func (p *Provider) ThatInt64(actual int64) *Int64 {
	return &Int64{p.logFacade, actual}
}

func (p *Provider) ThatUint64(actual uint64) *Uint64 {
	return &Uint64{p.logFacade, actual}
}

func (p *Provider) ThatFloat(actual float64) *Float {
	return &Float{p.logFacade, actual}
}

func (p *Provider) ThatComplex(actual complex128) *Complex {
	return &Complex{p.logFacade, actual}
}

func (p *Provider) ThatString(actual string) *String {
	return &String{p.logFacade, actual}
}

func (p *Provider) ThatBool(actual bool) *Bool {
	return &Bool{p.logFacade, actual}
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
