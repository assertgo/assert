package assert

import (
  "regexp"
  "strings"
)

type stringAssertImpl struct {
	logFacade *logFacade
	actual    string
}

func (assert *stringAssertImpl) isTrue(condition bool, format string, args ...interface{}) *stringAssertImpl {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *stringAssertImpl) IsEqualTo(expected string) StringAssert {
	return assert.isTrue(assert.actual == expected,
		"Expected <%s>, but was <%s>.", expected, assert.actual)
}

func (assert *stringAssertImpl) IsNotEqualTo(unexpected string) StringAssert {
	return assert.isTrue(assert.actual != unexpected,
		"Expected string not equal to <%s>, but was equal.", unexpected)
}

func (assert *stringAssertImpl) IsEmpty() StringAssert {
	return assert.isTrue(assert.actual == "",
		"Expected string to be empty, but was <%s>.", assert.actual)
}

func (assert *stringAssertImpl) IsNotEmpty() StringAssert {
	return assert.isTrue(assert.actual != "",
		"Expected string to not be empty, but was.",
	)
}

func (assert *stringAssertImpl) IsInSlice(expectedSlice []string) StringAssert {
	return assert.isTrue(stringIsInSlice(expectedSlice, assert.actual),
		"Expected string to be in slice <%v>, but wasn't.", expectedSlice)
}

func (assert *stringAssertImpl) IsNotInSlice(expectedSlice []string) StringAssert {
	return assert.isTrue(!stringIsInSlice(expectedSlice, assert.actual),
		"Expected string to not be in slice <%v>, but was.", expectedSlice)
}

func (assert *stringAssertImpl) Contains(substring string) StringAssert {
	return assert.isTrue(strings.Contains(assert.actual, substring),
		"Expected string <%s> to contain <%s>, but didn't.", assert.actual, substring)
}

func (assert *stringAssertImpl) DoesNotContain(substring string) StringAssert {
	return assert.isTrue(!strings.Contains(assert.actual, substring),
		"Expected string <%s> to not contain <%s>, but it did.", assert.actual, substring)
}

func (assert *stringAssertImpl) IsLowerCase() StringAssert {
	return assert.isTrue(strings.ToLower(assert.actual) == assert.actual,
		"Expected string <%s> to be lower case, but wasn't.", assert.actual)
}

func (assert *stringAssertImpl) IsNotLowerCase() StringAssert {
	return assert.isTrue(strings.ToLower(assert.actual) != assert.actual,
		"Expected string <%s> to not be lower case, but was.", assert.actual)
}

func (assert *stringAssertImpl) IsUpperCase() StringAssert {
	return assert.isTrue(strings.ToUpper(assert.actual) == assert.actual,
		"Expected string <%s> to be upper case, but wasn't.", assert.actual)
}

func (assert *stringAssertImpl) IsNotUpperCase() StringAssert {
	return assert.isTrue(strings.ToUpper(assert.actual) != assert.actual,
		"Expected string <%s> to not be upper case, but was.", assert.actual)
}

func (assert *stringAssertImpl) Matches(pattern string) StringAssert {
  matched, _ := regexp.MatchString(pattern, assert.actual)
	return assert.isTrue(matched,
		"Expected string <%s> to match <%s>, but didn't.", assert.actual, pattern)
}

func stringIsInSlice(slice []string, expectedString string) bool {
	for _, v := range slice {
		if v == expectedString {
			return true
		}
	}
	return false
}
