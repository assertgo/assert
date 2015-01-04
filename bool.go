package assert

type Bool struct {
	logFacade *logFacade
	actual    bool
}

func (a *Bool) IsFalse() *Bool {
	return a.isTrue(a.actual == false,
		"Expected <false>, but was <%t>.", a.actual)
}

func (a *Bool) IsTrue() *Bool {
	return a.isTrue(a.actual == true,
		"Expected <true>, but was <%t>.", a.actual)
}

func (a *Bool) IsSchrödingersCat() *Bool {
	return a.isTrue(a.actual == false && a.actual == true,
		"Expected Schrödinger's cat, but was <%t>.", a.actual)
}

func (a *Bool) isTrue(condition bool, format string, args ...interface{}) *Bool {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
