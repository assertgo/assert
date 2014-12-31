package assert

type ComplexAssert interface {
	IsZero() ComplexAssert
	IsNonZero() ComplexAssert
	// TODO
	// IsEqualTo(expected complex128) ComplexAssert
	// IsNotEqualTo(unexpected complex128) ComplexAssert
	// IsIn(elements ...complex128) ComplexAssert
	// IsNotIn(elements ...complex128) ComplexAssert
}
