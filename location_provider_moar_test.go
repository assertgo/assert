package assert

import "testing"

func TestLocationProviderCorrectFileName(t *testing.T) {
	assert := Setup(t)
	location := provideLocation()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_moar_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProviderCorrectFileName")
}

func TestCalledFromYetAnotherFunction(t *testing.T) {
	assert := Setup(t)
	location := anotherFunction()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestCalledFromYetAnotherFunction")
	assert.ThatInt(location.Line).IsEqualTo(28)
}
