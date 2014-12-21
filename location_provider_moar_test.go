package assert

import "testing"

func TestLocationProviderCorrectFileName(t *testing.T) {
	assert := Setup(t)
	location := provideLocation()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_moar_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProviderCorrectFileName")
}
