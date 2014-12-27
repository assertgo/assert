package assert

type AnyTypeAssert interface {
	IsEqualTo(expected interface{}) AnyTypeAssert
	IsNotEqualTo(unexpected interface{}) AnyTypeAssert
	IsNil() AnyTypeAssert
	IsNotNil() AnyTypeAssert
	AsBool() BoolAssert
	AsInt() IntAssert
	AsString() StringAssert
}
