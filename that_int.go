package assert

func ThatInt(t TestingT, actual int) IntAssert {
	return &intAssertImpl{t, actual}
}

type IntAssert interface {
	IsZero() IntAssert
	IsNonzero() IntAssert
	IsEqualTo(int) IntAssert
	IsNotEqualTo(int) IntAssert
	IsPositive() IntAssert
	IsNonpositive() IntAssert
	IsNegative() IntAssert
	IsNonnegative() IntAssert
	IsGreaterThan(int) IntAssert
	IsGreaterOrEqualTo(int) IntAssert
	IsLessThan(int) IntAssert
	IsLessOrEqualTo(int) IntAssert
	IsBetween(int, int) IntAssert
	IsNotBetween(int, int) IntAssert
	IsIn(...int) IntAssert
	IsNotIn(...int) IntAssert
}
