package assert

func ThatInt(t TestingT, actual int) IntAssert {
	return &intAssertImpl{t, actual}
}

type IntAssert interface {
	IsZero() IntAssert
	IsNonZero() IntAssert
	IsEqualTo(expected int) IntAssert
	IsNotEqualTo(unexpected int) IntAssert
	IsPositive() IntAssert
	IsNonPositive() IntAssert
	IsNegative() IntAssert
	IsNonNegative() IntAssert
	IsGreaterThan(less int) IntAssert
	IsGreaterOrEqualTo(lessOrEqual int) IntAssert
	IsLessThan(greater int) IntAssert
	IsLessOrEqualTo(greaterOrEqual int) IntAssert
	IsBetween(lowerBound, upperBound int) IntAssert
	IsNotBetween(lowerBound, upperBound int) IntAssert
	IsIn(elements ...int) IntAssert
	IsNotIn(elements ...int) IntAssert
	// TODO
	// IsEven() IntAssert
	// IsOdd() IntAssert
	// IsPrime() IntAssert
}
