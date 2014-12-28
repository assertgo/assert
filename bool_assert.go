package assert

type BoolAssert interface {
	IsFalse() BoolAssert
	IsTrue() BoolAssert
	IsSchrödingersCat() BoolAssert
}
