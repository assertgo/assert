package assert

import (
	"strconv"
	"strings"
)

type intAssertImpl struct {
	t      TestingT
	actual int
}

func (assert *intAssertImpl) isTrue(condition bool, format string, args ...interface{}) *intAssertImpl {
	logIfFalse(assert.t, condition, format, args...)
	return assert
}

func (assert *intAssertImpl) IsZero() IntAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%d>.", assert.actual)
}

func (assert *intAssertImpl) IsNonZero() IntAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}

func (assert *intAssertImpl) IsEqualTo(expected int) IntAssert {
	return assert.isTrue(assert.actual == expected,
		"Expected integer equal to <%d>, but was <%d>.", expected, assert.actual)
}

func (assert *intAssertImpl) IsNotEqualTo(unexpected int) IntAssert {
	return assert.isTrue(assert.actual != unexpected,
		"Expected integer not equal to <%d>, but was equal.", unexpected)
}

func (assert *intAssertImpl) IsPositive() IntAssert {
	return assert.isTrue(assert.actual > 0,
		"Expected positive integer, but was <%d>.", assert.actual)
}

func (assert *intAssertImpl) IsNonPositive() IntAssert {
	return assert.isTrue(assert.actual <= 0,
		"Expected nonpositive integer, but was <%d>.", assert.actual)
}

func (assert *intAssertImpl) IsNegative() IntAssert {
	return assert.isTrue(assert.actual < 0,
		"Expected negative integer, but was <%d>.", assert.actual)
}

func (assert *intAssertImpl) IsNonNegative() IntAssert {
	return assert.isTrue(assert.actual >= 0,
		"Expected nonnegative integer, but was <%d>.", assert.actual)
}

func (assert *intAssertImpl) IsGreaterThan(less int) IntAssert {
	return assert.isTrue(assert.actual > less,
		"Expected integer greater than <%d>, but was <%d>.", less, assert.actual)
}

func (assert *intAssertImpl) IsGreaterOrEqualTo(lessOrEqual int) IntAssert {
	return assert.isTrue(assert.actual >= lessOrEqual,
		"Expected integer greater or equal to <%d>, but was <%d>.", lessOrEqual, assert.actual)
}

func (assert *intAssertImpl) IsLessThan(greater int) IntAssert {
	return assert.isTrue(assert.actual < greater,
		"Expected integer lesser than <%d>, but was <%d>.", greater, assert.actual)
}

func (assert *intAssertImpl) IsLessOrEqualTo(greaterOrEqual int) IntAssert {
	return assert.isTrue(assert.actual <= greaterOrEqual,
		"Expected integer lesser or equal to <%d>, but was <%d>.", greaterOrEqual, assert.actual)
}

func (assert *intAssertImpl) IsBetween(lowerBound, upperBound int) IntAssert {
	return assert.isTrue(lowerBound <= assert.actual && assert.actual <= upperBound,
		"Expected integer to be between <%d, %d>, but was <%d>.", lowerBound, upperBound, assert.actual)
}

func (assert *intAssertImpl) IsNotBetween(lowerBound, upperBound int) IntAssert {
	return assert.isTrue(assert.actual < lowerBound || assert.actual > upperBound,
		"Expected integer not to be between <%d, %d>, but was <%d>.", lowerBound, upperBound, assert.actual)
}

func (assert *intAssertImpl) IsIn(elements ...int) IntAssert {
	return assert.isTrue(assert.isIn(elements...),
		"Expected integer to be in (%s), but was <%d>.", strings.Join(strSlice(elements...), ", "), assert.actual)
}

func (assert *intAssertImpl) IsNotIn(elements ...int) IntAssert {
	return assert.isTrue(!assert.isIn(elements...),
		"Expected integer not to be in (%s), but was <%d>.", strings.Join(strSlice(elements...), ", "), assert.actual)
}

func (assert *intAssertImpl) isIn(elements ...int) bool {
	for _, e := range elements {
		if assert.actual == e {
			return true
		}
	}
	return false
}

func strSlice(values ...int) []string {
	ret := make([]string, len(values))
	for i, v := range values {
		ret[i] = strconv.Itoa(v)
	}
	return ret
}
