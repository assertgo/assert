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
