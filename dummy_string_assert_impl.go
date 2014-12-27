package assert

type dummyStringAssertImpl struct {
}

func (assert *dummyStringAssertImpl) IsEqualTo(expected string) StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsNotEqualTo(unexpected string) StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsEmpty() StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsNotEmpty() StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsInSlice(expectedSlice []string) StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsNotInSlice(expectedSlice []string) StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) Contains(substring string) StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) DoesNotContain(substring string) StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsLowerCase() StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsNotLowerCase() StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsUpperCase() StringAssert {
	return assert
}

func (assert *dummyStringAssertImpl) IsNotUpperCase() StringAssert {
	return assert
}
