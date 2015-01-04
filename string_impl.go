package assert

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
