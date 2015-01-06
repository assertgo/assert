package assert

import "testing"

func TestLocationProviderCorrectFileName(t *testing.T) {
	assert := New(t)
	thisLineLocation := provideLocation(0)
	assert.That(thisLineLocation).IsEqualTo(
		&location{"TestLocationProviderCorrectFileName", "location_provider_moar_test.go", 7},
	)
}

func TestLocationProviderCalledFromAnotherFile(t *testing.T) {
	assert := New(t)
	insideAnotherFunctionLocation := anotherFunction()
	assert.That(insideAnotherFunctionLocation).IsEqualTo(
		&location{"TestLocationProviderCalledFromAnotherFile", "location_provider_test.go", 6},
	)
}
