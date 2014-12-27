package assert

type BoolAssert interface {
	IsFalse() BoolAssert
	IsTrue() BoolAssert
}
