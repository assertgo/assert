package assert

import (
	"bytes"
	"testing"
)

func TestThatIntIsZero(t *testing.T) {
	ThatInt(0).IsZero()
}

func TestThatIntIsZeroFails(t *testing.T) {
	defer _recover()
	ThatInt(1).IsZero()
	t.Fail()
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	b := &bytes.Buffer{}
	prod := writer
	writer = b
	defer func() {
		recover()
		if b.String() != "Expected <0>, but was <1>.\n" {
			t.Fail()
		}
		writer = prod
	}()
	ThatInt(1).IsZero()
}

func _recover() {
	recover()
}
