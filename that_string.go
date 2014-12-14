package assert

func ThatString(t TestingT, actual string) StringAssert {
	return &stringAssertImpl{t, actual}
}

type StringAssert interface {
	IsEqualTo(expected string) StringAssert
	IsNotEqualTo(expected string) StringAssert
}
