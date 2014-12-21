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
