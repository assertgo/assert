package assert

import "testing"

func anotherFunction() *location {
	return provideLocation(0)
}

func TestLocationProvider(t *testing.T) {
	assert := Setup(t)
	thisLineLocation := provideLocation(0)
	assert.That(thisLineLocation).IsEqualTo(
		&location{"TestLocationProvider", "location_provider_test.go", 11},
	)
}

func TestLocationProviderCalledFromAnotherFunction(t *testing.T) {
	assert := Setup(t)
	insideAnotherFunctionLocation := anotherFunction()
	assert.That(insideAnotherFunctionLocation).IsEqualTo(
		&location{"TestLocationProviderCalledFromAnotherFunction", "location_provider_test.go", 6},
	)
}
