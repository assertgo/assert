package assert

type TestingT interface {
	Fail()
}

func Setup(t TestingT) AssertProvider {
	return setupImpl(t, theLogger)
}

type AssertProvider interface {
	That(actual interface{}) AnyTypeAssert
	ThatInt(actual int) IntAssert
	ThatFloat(actual float64) FloatAssert
	ThatComplex(actual complex128) ComplexAssert
	ThatString(actual string) StringAssert
	ThatBool(actual bool) BoolAssert
	// TODO
	// ThatInt64(actual int64) Int64Assert
	// ThatUint64(actual uint64) Uint64Assert
}
