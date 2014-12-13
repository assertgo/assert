package assert

import "testing"

func TestThatStringIsEqualTo(t *testing.T) {
	ThatString("Alice has a cat.").IsEqualTo("Alice has a cat.")
}

func TestThatStringIsEqualToFails(t *testing.T) {
	defer _recover()
	ThatString("Alice has a cat.").IsEqualTo("Bob has a dog.")
	t.Fail()
}

func _recover() {
	recover()
}
