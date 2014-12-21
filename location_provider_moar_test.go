package assert

import "testing"

func TestLocationProviderCorrectFileName(t *testing.T) {
	assert := Setup(t)
	location := provideLocation(0)
	assert.ThatString(location.FileName).IsEqualTo("location_provider_moar_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProviderCorrectFileName")
	assert.ThatInt(location.Line).IsEqualTo(7)
}

func TestLocationProviderCalledFromAnotherFile(t *testing.T) {
	assert := Setup(t)
	location := anotherFunction()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProviderCalledFromAnotherFile")
	assert.ThatInt(location.Line).IsEqualTo(6)
}
