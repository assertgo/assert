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
