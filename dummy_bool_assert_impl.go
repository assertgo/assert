package assert

type dummyBoolAssertImpl struct {
}

func (assert *dummyBoolAssertImpl) IsFalse() BoolAssert {
	return assert
}

func (assert *dummyBoolAssertImpl) IsTrue() BoolAssert {
	return assert
}

func (assert *dummyBoolAssertImpl) IsSchrödingersCat() BoolAssert {
	return assert
}
