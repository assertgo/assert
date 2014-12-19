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
	IsNotInSlice(elements []string) StringAssert
	// TODO
	// Matches(pattern string) StringAssert
	// IsNotEmpty() StringAssert
	// DoesNotContain(substring string) StringAssert
	// IsIn(elements ...string) StringAssert
	// IsNotIn(elements ...string) StringAssert
	// HasPrefix(prefix string) StringAssert
	// HasSuffix(suffix string) StringAssert
	// IsLowerCase() StringAssert
	// IsUpperCase() StringAssert
	// IgnoringCase() StringAssert // new assert that doesn't care about case
	// Length() IntAssert // new assert to check on length of the string
}
