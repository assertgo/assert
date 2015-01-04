package assert

import "strconv"

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
