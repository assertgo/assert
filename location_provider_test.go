package assert

import "testing"

func anotherFunction() *location {
	return provideLocation(0)
}

func TestLocationProvider(t *testing.T) {
	assert := New(t)
	thisLineLocation := provideLocation(0)
	assert.That(thisLineLocation).IsEqualTo(
		&location{"TestLocationProvider", "location_provider_test.go", 11},
	)
}

func TestLocationProviderCalledFromAnotherFunction(t *testing.T) {
	assert := New(t)
	insideAnotherFunctionLocation := anotherFunction()
	assert.That(insideAnotherFunctionLocation).IsEqualTo(
		&location{"TestLocationProviderCalledFromAnotherFunction", "location_provider_test.go", 6},
	)
}

func TestLocationProviderPanicsWhenUsedIncorrectly(t *testing.T) {
	assert := New(t)
	defer func() {
		r := recover()
		assert.That(r).IsEqualTo("you shall not pass!")
	}()
	provideLocation(1)
	t.Fail()
}
