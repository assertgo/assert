package assert

import "testing"

func TestThatIntIsZero(t *testing.T) {
	ThatInt(0).IsZero()
}
