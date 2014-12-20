package assert

func ThatBool(t TestingT, actual bool) BoolAssert {
	return &boolAssertImpl{t, newLogger(), actual}
}

type BoolAssert interface {
	IsFalse() BoolAssert
	IsTrue() BoolAssert
}
