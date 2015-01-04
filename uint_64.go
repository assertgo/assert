package assert

type Uint64 struct {
	logFacade *logFacade
	actual    uint64
}

func (a *Uint64) IsZero() *Uint64 {
	return a.isTrue(a.actual == 0,
		"Expected zero, but was <%v>.", a.actual)
}

func (a *Uint64) IsNonZero() *Uint64 {
	return a.isTrue(a.actual != 0,
		"Expected nonzero, but was zero.")
}
