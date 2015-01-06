package assert

import (
	"fmt"
	"testing"
)

func New(t *testing.T) *Provider {
	return setupImpl(t, theLogger)
}

type Provider struct {
	logFacade *logFacade
}

func (p *Provider) That(actual interface{}) *AnyType {
	return &AnyType{p.logFacade, actual}
}

func (p *Provider) ThatBool(actual bool) *Bool {
	return &Bool{p.logFacade, actual}
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

type testingT interface {
	Fail()
}

type logFacade struct {
	t      testingT
	logger errorLogger
}

func setupImpl(t testingT, logger errorLogger) *Provider {
	return &Provider{&logFacade{t, logger}}
}

func logIfFalse(lf *logFacade, condition bool, format string, args ...interface{}) {
	if lf != nil && !condition {
		location := provideLocation(3)
		message := fmt.Sprintf(format, args...)
		lf.logger.Log(location, message)
		lf.t.Fail()
	}
}
