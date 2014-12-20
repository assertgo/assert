package assert

import "os"

func ThatBool(t TestingT, actual bool) BoolAssert {
	return &boolAssertImpl{t, &errorLoggerImpl{os.Stderr}, actual}
}

type BoolAssert interface {
	IsFalse() BoolAssert
	IsTrue() BoolAssert
}
