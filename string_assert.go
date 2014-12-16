package assert

func ThatString(t TestingT, actual string) StringAssert {
	return &stringAssertImpl{t, actual}
}

type StringAssert interface {
	IsEqualTo(expected string) StringAssert
	IsNotEqualTo(unexpected string) StringAssert
	IsEmpty() StringAssert
	IsInSlice(expectedSlice []string) StringAssert
	Contains(substring string) StringAssert
	// TODO
	// Matches(pattern string) StringAssert
}
