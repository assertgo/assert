package assert

type AnyTypeAssert interface {
	IsEqualTo(expected interface{}) AnyTypeAssert
	IsNotEqualTo(unexpected interface{}) AnyTypeAssert
	IsNil() AnyTypeAssert
	IsNotNil() AnyTypeAssert
	AsBool() BoolAssert
	AsInt() IntAssert
	AsString() StringAssert
	// TODO
	// AsInt64() Int64Assert
	// AsUint64() Uint64Assert
	// AsFloat() FloatAssert
	// AsComplex() ComplexAssert
}
