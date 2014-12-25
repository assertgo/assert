package assert

type dummyIntAssertImpl struct {
}

func (assert *dummyIntAssertImpl) IsZero() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNonZero() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsEqualTo(expected int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNotEqualTo(unexpected int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsPositive() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNonPositive() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNegative() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNonNegative() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsGreaterThan(less int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsGreaterOrEqualTo(lessOrEqual int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsLessThan(greater int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsLessOrEqualTo(greaterOrEqual int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsBetween(lowerBound, upperBound int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNotBetween(lowerBound, upperBound int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsIn(elements ...int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNotIn(elements ...int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsEven() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsOdd() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsDivisibleBy(divisor int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNotDivisibleBy(divisor int) IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsPrime() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsNotPrime() IntAssert {
	return assert
}

func (assert *dummyIntAssertImpl) IsAnswerToTheUltimateQuestionOfLife() IntAssert {
	return assert
}
