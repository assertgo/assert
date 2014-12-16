package assert

func ThatString(t TestingT, actual string) StringAssert {
	return &stringAssertImpl{t, actual}
}

type StringAssert interface {
	IsEqualTo(expected string) StringAssert
	IsNotEqualTo(unexpected string) StringAssert
	IsEmpty() StringAssert
	IsInSlice(expectedSlice []string) StringAssert
	// TODO
	// Contains(substring string) StringAssert
	// Matches(pattern string) StringAssert
}
