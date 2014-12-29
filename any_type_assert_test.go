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
