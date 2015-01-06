package assert

import (
	"strconv"
	"strings"
)

type Int struct {
	logFacade *logFacade
	actual    int
}

func (a *Int) IsZero() *Int {
	return a.isTrue(a.actual == 0,
		"Expected zero, but was <%d>.", a.actual)
}

func (a *Int) IsNonZero() *Int {
	return a.isTrue(a.actual != 0,
		"Expected nonzero, but was zero.")
}

func (a *Int) IsEqualTo(expected int) *Int {
	return a.isTrue(a.actual == expected,
		"Expected integer equal to <%d>, but was <%d>.", expected, a.actual)
}

func (a *Int) IsNotEqualTo(unexpected int) *Int {
	return a.isTrue(a.actual != unexpected,
		"Expected integer not equal to <%d>, but was equal.", unexpected)
}

func (a *Int) IsPositive() *Int {
	return a.isTrue(a.actual > 0,
		"Expected positive integer, but was <%d>.", a.actual)
}

func (a *Int) IsNonPositive() *Int {
	return a.isTrue(a.actual <= 0,
		"Expected nonpositive integer, but was <%d>.", a.actual)
}

func (a *Int) IsNegative() *Int {
	return a.isTrue(a.actual < 0,
		"Expected negative integer, but was <%d>.", a.actual)
}

func (a *Int) IsNonNegative() *Int {
	return a.isTrue(a.actual >= 0,
		"Expected nonnegative integer, but was <%d>.", a.actual)
}

func (a *Int) IsGreaterThan(less int) *Int {
	return a.isTrue(a.actual > less,
		"Expected integer greater than <%d>, but was <%d>.", less, a.actual)
}

func (a *Int) IsGreaterOrEqualTo(lessOrEqual int) *Int {
	return a.isTrue(a.actual >= lessOrEqual,
		"Expected integer greater or equal to <%d>, but was <%d>.", lessOrEqual, a.actual)
}

func (a *Int) IsLessThan(greater int) *Int {
	return a.isTrue(a.actual < greater,
		"Expected integer lesser than <%d>, but was <%d>.", greater, a.actual)
}

func (a *Int) IsLessOrEqualTo(greaterOrEqual int) *Int {
	return a.isTrue(a.actual <= greaterOrEqual,
		"Expected integer lesser or equal to <%d>, but was <%d>.", greaterOrEqual, a.actual)
}

func (a *Int) IsBetween(lowerBound, upperBound int) *Int {
	return a.isTrue(lowerBound <= a.actual && a.actual <= upperBound,
		"Expected integer to be between <%d, %d>, but was <%d>.", lowerBound, upperBound, a.actual)
}

func (a *Int) IsNotBetween(lowerBound, upperBound int) *Int {
	return a.isTrue(a.actual < lowerBound || a.actual > upperBound,
		"Expected integer not to be between <%d, %d>, but was <%d>.", lowerBound, upperBound, a.actual)
}

func (a *Int) IsIn(elements ...int) *Int {
	return a.isTrue(a.isIn(elements...),
		"Expected integer to be in (%s), but was <%d>.", strings.Join(strSlice(elements...), ", "), a.actual)
}

func (a *Int) IsNotIn(elements ...int) *Int {
	return a.isTrue(!a.isIn(elements...),
		"Expected integer not to be in (%s), but was <%d>.", strings.Join(strSlice(elements...), ", "), a.actual)
}

func (a *Int) IsEven() *Int {
	return a.isTrue(a.actual%2 == 0,
		"Expected integer to be even, but was <%d>.", a.actual)
}

func (a *Int) IsOdd() *Int {
	return a.isTrue(a.actual%2 != 0,
		"Expected integer to be odd, but was <%d>.", a.actual)
}

func (a *Int) IsDivisibleBy(divisor int) *Int {
	return a.isTrue(a.actual%divisor == 0,
		"Expected integer to be divisible by <%d>, but was <%d>.", divisor, a.actual)
}

func (a *Int) IsNotDivisibleBy(divisor int) *Int {
	return a.isTrue(a.actual%divisor != 0,
		"Expected integer not to be divisible by <%d>, but was <%d>.", divisor, a.actual)
}

func (a *Int) IsPrime() *Int {
	return a.isTrue(isPrime(a.actual),
		"Expected integer to be prime, but was <%d>.", a.actual)
}

func (a *Int) IsNotPrime() *Int {
	return a.isTrue(!isPrime(a.actual),
		"Expected integer not to be prime, but was <%d>.", a.actual)
}

func (a *Int) IsAnswerToTheUltimateQuestionOfLife() *Int {
	return a.isTrue(a.actual == 42,
		"Expected answer to the ultimate question of life, but was <%d>.", a.actual)
}

func (a *Int) isTrue(condition bool, format string, args ...interface{}) *Int {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}

func (a *Int) isIn(elements ...int) bool {
	for _, e := range elements {
		if a.actual == e {
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
