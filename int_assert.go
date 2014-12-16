package assert

func ThatInt(t TestingT, actual int) IntAssert {
	return &intAssertImpl{t, actual}
}

type IntAssert interface {
	IsZero() IntAssert
	IsNonZero() IntAssert
	IsEqualTo(int) IntAssert
	IsNotEqualTo(int) IntAssert
	IsPositive() IntAssert
	IsNonPositive() IntAssert
	IsNegative() IntAssert
	IsNonNegative() IntAssert
	IsGreaterThan(int) IntAssert
	IsGreaterOrEqualTo(int) IntAssert
	IsLessThan(int) IntAssert
	IsLessOrEqualTo(int) IntAssert
	IsBetween(int, int) IntAssert
	IsNotBetween(int, int) IntAssert
	IsIn(...int) IntAssert
	IsNotIn(...int) IntAssert
	// TODO
	// IsEven() IntAssert
	// IsOdd() IntAssert
	// IsPrime() IntAssert
}
