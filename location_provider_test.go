package assert

import "testing"

func TestLocationProvider(t *testing.T) {
	assert := Setup(t)
	location := provideLocation()
	assert.ThatString(location.FileName).IsEqualTo("location_provider_test.go")
	assert.ThatString(location.Test).IsEqualTo("TestLocationProvider")
}
