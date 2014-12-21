package assert

import "testing"

func anotherFunction() *location {
	return provideLocation(0)
}

func TestLocationProvider(t *testing.T) {
	assert := Setup(t)
	location := provideLocation(0)
	assert.ThatString(location.FileName).IsEqualTo("location_provider_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProvider")
	assert.ThatInt(location.Line).IsEqualTo(11)
}

func TestLocationProviderCalledFromAnotherFunction(t *testing.T) {
	assert := Setup(t)
	location := anotherFunction()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProviderCalledFromAnotherFunction")
	assert.ThatInt(location.Line).IsEqualTo(6)
}
