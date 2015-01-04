package assert

import (
	"strconv"
	"strings"
)

type IntAssert struct {
	logFacade *logFacade
	actual    int
}

func (assert *IntAssert) isTrue(condition bool, format string, args ...interface{}) *IntAssert {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *IntAssert) IsZero() *IntAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonZero() *IntAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}

func (assert *IntAssert) IsEqualTo(expected int) *IntAssert {
	return assert.isTrue(assert.actual == expected,
		"Expected integer equal to <%d>, but was <%d>.", expected, assert.actual)
}

func (assert *IntAssert) IsNotEqualTo(unexpected int) *IntAssert {
	return assert.isTrue(assert.actual != unexpected,
		"Expected integer not equal to <%d>, but was equal.", unexpected)
}

func (assert *IntAssert) IsPositive() *IntAssert {
	return assert.isTrue(assert.actual > 0,
		"Expected positive integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonPositive() *IntAssert {
	return assert.isTrue(assert.actual <= 0,
		"Expected nonpositive integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNegative() *IntAssert {
	return assert.isTrue(assert.actual < 0,
		"Expected negative integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonNegative() *IntAssert {
	return assert.isTrue(assert.actual >= 0,
		"Expected nonnegative integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsGreaterThan(less int) *IntAssert {
	return assert.isTrue(assert.actual > less,
		"Expected integer greater than <%d>, but was <%d>.", less, assert.actual)
}

func (assert *IntAssert) IsGreaterOrEqualTo(lessOrEqual int) *IntAssert {
	return assert.isTrue(assert.actual >= lessOrEqual,
		"Expected integer greater or equal to <%d>, but was <%d>.", lessOrEqual, assert.actual)
}

func (assert *IntAssert) IsLessThan(greater int) *IntAssert {
	return assert.isTrue(assert.actual < greater,
		"Expected integer lesser than <%d>, but was <%d>.", greater, assert.actual)
}

func (assert *IntAssert) IsLessOrEqualTo(greaterOrEqual int) *IntAssert {
	return assert.isTrue(assert.actual <= greaterOrEqual,
		"Expected integer lesser or equal to <%d>, but was <%d>.", greaterOrEqual, assert.actual)
}

func (assert *IntAssert) IsBetween(lowerBound, upperBound int) *IntAssert {
	return assert.isTrue(lowerBound <= assert.actual && assert.actual <= upperBound,
		"Expected integer to be between <%d, %d>, but was <%d>.", lowerBound, upperBound, assert.actual)
}

func (assert *IntAssert) IsNotBetween(lowerBound, upperBound int) *IntAssert {
	return assert.isTrue(assert.actual < lowerBound || assert.actual > upperBound,
		"Expected integer not to be between <%d, %d>, but was <%d>.", lowerBound, upperBound, assert.actual)
}

func (assert *IntAssert) IsIn(elements ...int) *IntAssert {
	return assert.isTrue(assert.isIn(elements...),
		"Expected integer to be in (%s), but was <%d>.", strings.Join(strSlice(elements...), ", "), assert.actual)
}

func (assert *IntAssert) IsNotIn(elements ...int) *IntAssert {
	return assert.isTrue(!assert.isIn(elements...),
		"Expected integer not to be in (%s), but was <%d>.", strings.Join(strSlice(elements...), ", "), assert.actual)
}

func (assert *IntAssert) IsEven() *IntAssert {
	return assert.isTrue(assert.actual%2 == 0,
		"Expected integer to be even, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsOdd() *IntAssert {
	return assert.isTrue(assert.actual%2 != 0,
		"Expected integer to be odd, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsDivisibleBy(divisor int) *IntAssert {
	return assert.isTrue(assert.actual%divisor == 0,
		"Expected integer to be divisible by <%d>, but was <%d>.", divisor, assert.actual)
}

func (assert *IntAssert) IsNotDivisibleBy(divisor int) *IntAssert {
	return assert.isTrue(assert.actual%divisor != 0,
		"Expected integer not to be divisible by <%d>, but was <%d>.", divisor, assert.actual)
}

func (assert *IntAssert) IsPrime() *IntAssert {
	return assert.isTrue(isPrime(assert.actual),
		"Expected integer to be prime, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNotPrime() *IntAssert {
	return assert.isTrue(!isPrime(assert.actual),
		"Expected integer not to be prime, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsAnswerToTheUltimateQuestionOfLife() *IntAssert {
	return assert.isTrue(assert.actual == 42,
		"Expected answer to the ultimate question of life, but was <%d>.", assert.actual)
}

func (assert *IntAssert) isIn(elements ...int) bool {
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
