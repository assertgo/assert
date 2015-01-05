package assert

import (
	"regexp"
	"strings"
)

type String struct {
	logFacade *logFacade
	actual    string
}

func (a *String) IsEqualTo(expected string) *String {
	return a.isTrue(a.actual == expected,
		"Expected <%s>, but was <%s>.", expected, a.actual)
}

func (a *String) IsNotEqualTo(unexpected string) *String {
	return a.isTrue(a.actual != unexpected,
		"Expected string not equal to <%s>, but was equal.", unexpected)
}

func (a *String) IsEmpty() *String {
	return a.isTrue(a.actual == "",
		"Expected string to be empty, but was <%s>.", a.actual)
}

func (a *String) IsNotEmpty() *String {
	return a.isTrue(a.actual != "",
		"Expected string to not be empty, but was.",
	)
}

func (a *String) IsInSlice(expectedSlice []string) *String {
	return a.isTrue(stringIsInSlice(expectedSlice, a.actual),
		"Expected string to be in slice <%v>, but wasn't.", expectedSlice)
}

func (a *String) IsNotInSlice(expectedSlice []string) *String {
	return a.isTrue(!stringIsInSlice(expectedSlice, a.actual),
		"Expected string to not be in slice <%v>, but was.", expectedSlice)
}

func (a *String) Contains(substring string) *String {
	return a.isTrue(strings.Contains(a.actual, substring),
		"Expected string <%s> to contain <%s>, but didn't.", a.actual, substring)
}

func (a *String) DoesNotContain(substring string) *String {
	return a.isTrue(!strings.Contains(a.actual, substring),
		"Expected string <%s> to not contain <%s>, but it did.", a.actual, substring)
}

func (a *String) IsLowerCase() *String {
	return a.isTrue(strings.ToLower(a.actual) == a.actual,
		"Expected string <%s> to be lower case, but wasn't.", a.actual)
}

func (a *String) IsNotLowerCase() *String {
	return a.isTrue(strings.ToLower(a.actual) != a.actual,
		"Expected string <%s> to not be lower case, but was.", a.actual)
}

func (a *String) IsUpperCase() *String {
	return a.isTrue(strings.ToUpper(a.actual) == a.actual,
		"Expected string <%s> to be upper case, but wasn't.", a.actual)
}

func (a *String) IsNotUpperCase() *String {
	return a.isTrue(strings.ToUpper(a.actual) != a.actual,
		"Expected string <%s> to not be upper case, but was.", a.actual)
}

func (a *String) Matches(pattern string) *String {
	matched, _ := regexp.MatchString(pattern, a.actual)
	return a.isTrue(matched,
		"Expected string <%s> to match <%s>, but didn't.", a.actual, pattern)
}

func (a *String) isTrue(condition bool, format string, args ...interface{}) *String {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}

func stringIsInSlice(slice []string, expectedString string) bool {
	for _, v := range slice {
		if v == expectedString {
			return true
		}
	}
	return false
}
