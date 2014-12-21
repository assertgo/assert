package assert

type AnyTypeAssert interface {
	IsEqualTo(expected interface{}) AnyTypeAssert
}
