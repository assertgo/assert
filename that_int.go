package assert

import (
	"strconv"
	"strings"
)

func ThatInt(t TestingT, actual int) *IntAssert {
	return &IntAssert{t, actual}
}

type IntAssert struct {
	t      TestingT
	actual int
}

func (assert *IntAssert) isTrue(condition bool, format string, args ...interface{}) *IntAssert {
	if !condition {
		assert.t.Errorf(format, args...)
	}
	return assert
}

func (assert *IntAssert) IsZero() *IntAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonzero() *IntAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}

func (assert *IntAssert) IsEqualTo(expected int) *IntAssert {
	return assert.isTrue(assert.actual == expected,
		"Expected integer equal to <%d>, but was <%d>.", expected, assert.actual)
}

func (assert *IntAssert) IsNotEqualTo(nonexpected int) *IntAssert {
	return assert.isTrue(assert.actual != nonexpected,
		"Expected integer not equal to <%d>, but was equal.", nonexpected)
}

func (assert *IntAssert) IsPositive() *IntAssert {
	return assert.isTrue(assert.actual > 0,
		"Expected positive integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonpositive() *IntAssert {
	return assert.isTrue(assert.actual <= 0,
		"Expected nonpositive integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNegative() *IntAssert {
	return assert.isTrue(assert.actual < 0,
		"Expected negative integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsNonnegative() *IntAssert {
	return assert.isTrue(assert.actual >= 0,
		"Expected nonnegative integer, but was <%d>.", assert.actual)
}

func (assert *IntAssert) IsGreaterThan(lesser int) *IntAssert {
	return assert.isTrue(assert.actual > lesser,
		"Expected integer greater than <%d>, but was <%d>.", lesser, assert.actual)
}

func (assert *IntAssert) IsGreaterOrEqualTo(lesserOrEqual int) *IntAssert {
	return assert.isTrue(assert.actual >= lesserOrEqual,
		"Expected integer greater or equal to <%d>, but was <%d>.", lesserOrEqual, assert.actual)
}

func (assert *IntAssert) IsLessThan(greater int) *IntAssert {
	return assert.isTrue(assert.actual < greater,
		"Expected integer lesser than <%d>, but was <%d>.", greater, assert.actual)
}

func (assert *IntAssert) IsLessOrEqualTo(greaterOrEqual int) *IntAssert {
	return assert.isTrue(assert.actual <= greaterOrEqual,
		"Expected integer lesser or equal to <%d>, but was <%d>.", greaterOrEqual, assert.actual)
}

func (assert *IntAssert) IsBetween(lesserOrEqual, greaterOrEqual int) *IntAssert {
	return assert.isTrue(lesserOrEqual <= assert.actual && assert.actual <= greaterOrEqual,
		"Expected integer to be between <%d, %d>, but was <%d>.", lesserOrEqual, greaterOrEqual, assert.actual)
}

func (assert *IntAssert) IsNotBetween(nonlesser, nongreater int) *IntAssert {
	return assert.isTrue(assert.actual < nonlesser || assert.actual > nongreater,
		"Expected integer not to be between <%d, %d>, but was <%d>.", nonlesser, nongreater, assert.actual)
}

func (assert *IntAssert) IsIn(expected ...int) *IntAssert {
	return assert.isTrue(assert.isIn(expected...),
		"Expected integer to be in (%s), but was <%d>.", strings.Join(strSlice(expected...), ", "), assert.actual)
}

func (assert *IntAssert) isIn(expected ...int) bool {
	for _, v := range expected {
		if assert.actual == v {
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
