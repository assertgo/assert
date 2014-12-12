package assert

import "testing"

func TestThatIntIsZero(t *testing.T) {
	ThatInt(0).IsZero()
}

func TestThatIntIsZeroFails(t *testing.T) {
	defer _recover()
	ThatInt(1).IsZero()
	t.Fail()
}

func _recover() {
	recover()
}
