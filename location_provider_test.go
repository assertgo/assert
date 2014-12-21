package assert

import "testing"

func TestLocationProvider(t *testing.T) {
	assert := Setup(t)
	location := provideLocation()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProvider")
	assert.ThatInt(location.Line).IsEqualTo(7)
}

func TestLocationProviderCase2(t *testing.T) {
	assert := Setup(t)
	location := provideLocation()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProviderCase2")
}

func TestCalledFromAnotherFunction(t *testing.T) {
	assert := Setup(t)
	location := anotherFunction()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestCalledFromAnotherFunction")
}

func anotherFunction() location {
	return provideLocation()
}
