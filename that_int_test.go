package assert

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestThatIntIsZero(t *testing.T) {
	ThatInt(0).IsZero()
}

func TestThatIntIsZeroFails(t *testing.T) {
	recoverAndRestore := mockWriter(ioutil.Discard)
	defer recoverAndRestore()
	ThatInt(1).IsZero()
	t.Fail()
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	assert := func() {
		if buffer.String() != "Expected <0>, but was <6>.\n" {
			t.Error("Expected <Expected <0>, but was <6>.\n>, but was <" + buffer.String() + ">.")
		}
	}
	recoverAndRestore := mockWriter(buffer, assert)
	defer recoverAndRestore()
	ThatInt(6).IsZero()
}

func TestThatIntIsEqualTo(t *testing.T) {
	ThatInt(1).IsEqualTo(1)
}

func TestThatIntIsEqualToFails(t *testing.T) {
	recoverAndRestore := mockWriter(ioutil.Discard)
	defer recoverAndRestore()
	ThatInt(1).IsEqualTo(2)
	t.Fail()
}

func ThestThatIntIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	assert := func() {
		if buffer.String() != "Expected <1>, but was <2>.\n" {
			t.Error("Expected <Expected <1>, but was <2>.\n>, but was <" + buffer.String() + ">.")
		}
	}
	recoverAndRestore := mockWriter(buffer, assert)
	defer recoverAndRestore()
	ThatInt(1).IsEqualTo(2)
}

func mockWriter(writerDouble io.Writer, asserts ...func()) func() {
	writerBackup := writer
	writer = writerDouble
	return func() {
		recover()
		callAsserts(asserts...)
		writer = writerBackup
	}
}

func callAsserts(asserts ...func()) {
	for _, assert := range asserts {
		assert()
	}
}
