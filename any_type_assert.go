package assert

type AnyTypeAssert interface {
	IsEqualTo(expected interface{}) AnyTypeAssert
	// TODO
	// IsNotEqualTo(unexpected interface{}) AnyTypeAssert
	// IsNil() AnyTypeAssert
	// IsNotNil() AnyTypeAssert
	// AsInt() IntAssert // new assert that treats actual value as int, reports failure if cannot convert
	// AsString() StringAssert
}
