package assert

import (
	"os"
)

func Setup(t TestingT) AssertProvider {
	return &assertProviderImpl{t, &errorLoggerImpl{os.Stderr}}
}

type AssertProvider interface {
	ThatInt(actual int) IntAssert
	ThatString(actual string) StringAssert
	ThatBool(actual bool) BoolAssert
}
