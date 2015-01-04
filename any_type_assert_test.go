package assert

import "testing"

func TestThatIsEqualToHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(&location{"TestFoo", "foo_test.go", 66}).
		IsEqualTo(&location{"TestFoo", "foo_test.go", 66}).
		IsEqualTo(&location{"TestFoo", "foo_test.go", 66})
	mockT.HasNoErrors()
}

func TestThatIsEqualToHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(location{"TestFoo", "foo_test.go", 66}).IsEqualTo(location{"TestBar", "bar_test.go", 66})
	assert.That(&location{"TestFoo", "foo_test.go", 66}).IsEqualTo(&location{"TestFoo", "foo_test.go", 33})
	mockT.HasErrorMessages(
		"Expected <{TestBar bar_test.go 66}>, but was <{TestFoo foo_test.go 66}>.",
		"Expected <&{TestFoo foo_test.go 33}>, but was <&{TestFoo foo_test.go 66}>.",
	)
}

func TestThatIsNotEqualToHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(location{"TestFoo", "foo_test.go", 66}).IsNotEqualTo(location{"TestBar", "bar_test.go", 66})
	assert.That(&location{"TestFoo", "foo_test.go", 66}).IsNotEqualTo(&location{"TestFoo", "foo_test.go", 33})
	mockT.HasNoErrors()
}

func TestThatIsNotEqualToHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(location{"TestFoo", "foo_test.go", 66}).
		IsNotEqualTo(location{"TestFoo", "foo_test.go", 66}).
		IsNotEqualTo(location{"TestFoo", "foo_test.go", 66})
	mockT.HasErrorMessages(
		"Expected value not equal to <{TestFoo foo_test.go 66}>, but was equal.",
		"Expected value not equal to <{TestFoo foo_test.go 66}>, but was equal.",
	)
}

func TestThatIsNilHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(nil).IsNil().IsNil()
	mockT.HasNoErrors()
}

func TestThatIsNilHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("nil").IsNil()
	assert.That(location{"TestFoo", "foo_test.go", 66}).IsNil()
	mockT.HasErrorMessages(
		"Expected value to be nil, but was <nil>.",
		"Expected value to be nil, but was <{TestFoo foo_test.go 66}>.",
	)
}

func TestThatIsNotNilHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("").IsNotNil()
	assert.That(0).IsNotNil()
	assert.That(location{"TestFoo", "foo_test.go", 66}).IsNotNil()
	assert.That(&location{"TestFoo", "foo_test.go", 66}).IsNotNil()
	mockT.HasNoErrors()
}

func TestThatIsNotNilHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(nil).IsNotNil().IsNotNil()
	mockT.HasErrorMessages(
		"Expected value not to be nil, but was.",
		"Expected value not to be nil, but was.",
	)
}

func TestThatAsBoolHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(false).AsBool().IsFalse()
	assert.That(true).AsBool().IsTrue()
	mockT.HasNoErrors()
}

func TestThatAsBoolHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("NO!").AsBool()
	assert.That(location{"TestFoo", "foo_test.go", 66}).AsBool()
	mockT.HasErrorMessages(
		"Cannot convert <NO!> of type <string> to <bool>.",
		"Cannot convert <{TestFoo foo_test.go 66}> of type <assert.location> to <bool>.",
	)
}

func TestThatAsIntHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(0).AsInt().IsZero()
	assert.That(6).AsInt().IsEqualTo(6)
	mockT.HasNoErrors()
}

func TestThatAsIntHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("NO!").AsInt()
	assert.That(&location{"TestFoo", "foo_test.go", 66}).AsInt()
	mockT.HasErrorMessages(
		"Cannot convert <NO!> of type <string> to <int>.",
		"Cannot convert <&{TestFoo foo_test.go 66}> of type <*assert.location> to <int>.",
	)
}

func TestThatAsInt64HasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(int64(0)).AsInt64().IsZero()
	assert.That(int64(7)).AsInt64().IsNonZero()
	mockT.HasNoErrors()
}

func TestThatAsInt64HasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("NO!").AsInt64()
	assert.That(&location{"TestFoo", "foo_test.go", 66}).AsInt64()
	mockT.HasErrorMessages(
		"Cannot convert <NO!> of type <string> to <int64>.",
		"Cannot convert <&{TestFoo foo_test.go 66}> of type <*assert.location> to <int64>.",
	)
}

func TestThatAsInt64CanAcceptOtherTypes(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(int8(-111)).AsInt64()
	assert.That(uint8(111)).AsInt64()
	assert.That(int16(-30000)).AsInt64()
	assert.That(uint16(65535)).AsInt64()
	assert.That(int32(-10000000)).AsInt64()
	assert.That(uint32(10000000)).AsInt64()
	assert.That(2147483647).AsInt64()
	mockT.HasNoErrors()
}

func TestThatAsUint64HasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(uint64(0)).AsUint64().IsZero()
	assert.That(uint64(7)).AsUint64().IsNonZero()
	mockT.HasNoErrors()
}

func TestThatAsUint64HasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("NO!").AsUint64()
	assert.That(&location{"TestFoo", "foo_test.go", 66}).AsUint64()
	mockT.HasErrorMessages(
		"Cannot convert <NO!> of type <string> to <uint64>.",
		"Cannot convert <&{TestFoo foo_test.go 66}> of type <*assert.location> to <uint64>.",
	)
}

func TestThatAsUint64CanAcceptOtherTypes(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(uint8(111)).AsUint64()
	assert.That(uint16(65535)).AsUint64()
	assert.That(uint32(10000000)).AsUint64()
	assert.That(uint(2147483648)).AsUint64()
	mockT.HasNoErrors()
}

func TestThatAsFloatHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(0.0).AsFloat().IsZero()
	assert.That(float64(-1)).AsFloat().IsNonZero()
	mockT.HasNoErrors()
}

func TestThatAsFloatHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("NO!").AsFloat()
	assert.That(&location{"TestFoo", "foo_test.go", 66}).AsFloat()
	mockT.HasErrorMessages(
		"Cannot convert <NO!> of type <string> to <float64>.",
		"Cannot convert <&{TestFoo foo_test.go 66}> of type <*assert.location> to <float64>.",
	)
}

func TestThatAsFloatCanAcceptOtherTypes(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(float32(-44.44)).AsFloat()
	assert.That(int8(-111)).AsFloat()
	assert.That(int16(-32768)).AsFloat()
	assert.That(int32(-10000000)).AsFloat()
	assert.That(uint8(111)).AsFloat()
	assert.That(uint16(65535)).AsFloat()
	assert.That(uint32(10000000)).AsFloat()
	mockT.HasNoErrors()
}

func TestThatAsComplexHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(complex(0, 0)).AsComplex().IsZero()
	assert.That(-2i).AsComplex().IsNonZero()
	mockT.HasNoErrors()
}

func TestThatAsComplexHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("NO!").AsComplex()
	assert.That(&location{"TestFoo", "foo_test.go", 66}).AsComplex()
	mockT.HasErrorMessages(
		"Cannot convert <NO!> of type <string> to <complex128>.",
		"Cannot convert <&{TestFoo foo_test.go 66}> of type <*assert.location> to <complex128>.",
	)
}

func TestThatAsComplexCanAcceptOtherTypes(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(complex(float32(1.1), float32(-1.1))).AsComplex()
	assert.That(float32(-44.44)).AsComplex()
	assert.That(float64(555.555)).AsComplex()
	assert.That(int8(-111)).AsComplex()
	assert.That(int16(-32768)).AsComplex()
	assert.That(int32(-10000000)).AsComplex()
	assert.That(uint8(111)).AsComplex()
	assert.That(uint16(65535)).AsComplex()
	assert.That(uint32(10000000)).AsComplex()
	mockT.HasNoErrors()
}

func TestThatAsStringHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That("").AsString().IsEmpty()
	assert.That("abcd").AsString().Contains("bc")
	mockT.HasNoErrors()
}

func TestThatAsStringHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(777).AsString()
	assert.That(&location{"TestFoo", "foo_test.go", 66}).AsString()
	mockT.HasErrorMessages(
		"Cannot convert <777> of type <int> to <string>.",
		"Cannot convert <&{TestFoo foo_test.go 66}> of type <*assert.location> to <string>.",
	)
}

func TestThatAsIntCanAcceptOtherTypes(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(int8(-111)).AsInt()
	assert.That(uint8(111)).AsInt()
	assert.That(int16(-30000)).AsInt()
	assert.That(uint16(65535)).AsInt()
	assert.That(int32(10000000)).AsInt()
	assert.That(byte(42)).AsInt()
	assert.That(rune('A')).AsInt()
	mockT.HasNoErrors()
}

type testCustomBool bool
type testCustomInt int
type testCustomString string

func TestThatConvertionCanAcceptCustomTypes(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(testCustomBool(true)).AsBool()
	assert.That(testCustomInt(6)).AsInt()
	assert.That(testCustomString("custom string")).AsString()
	mockT.HasNoErrors()
}
