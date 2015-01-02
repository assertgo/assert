package assert

type StringAssert interface {
	IsEqualTo(expected string) StringAssert
	IsNotEqualTo(unexpected string) StringAssert
	IsEmpty() StringAssert
	IsNotEmpty() StringAssert
	IsInSlice(expectedSlice []string) StringAssert
	IsNotInSlice(elements []string) StringAssert
	Contains(substring string) StringAssert
	DoesNotContain(substring string) StringAssert
	IsLowerCase() StringAssert
	IsNotLowerCase() StringAssert
	IsUpperCase() StringAssert
	IsNotUpperCase() StringAssert
  Matches(pattern string) StringAssert
	// TODO
	// IsIn(elements ...string) StringAssert
	// IsNotIn(elements ...string) StringAssert
	// HasPrefix(prefix string) StringAssert
	// HasSuffix(suffix string) StringAssert
	// IgnoringCase() StringAssert // new assert that doesn't care about case
	// Length() IntAssert // new assert to check on length of the string
	// LevenshteinDistance(other string) IntAssert // new assert to check in distance between words
}
