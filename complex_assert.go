package assert

type ComplexAssert interface {
	IsZero() ComplexAssert
	// TODO
	// IsNonZero() ComplexAssert
	// IsEqualTo(expected complex128) ComplexAssert
	// IsNotEqualTo(unexpected complex128) ComplexAssert
	// IsIn(elements ...complex128) ComplexAssert
	// IsNotIn(elements ...complex128) ComplexAssert
}
