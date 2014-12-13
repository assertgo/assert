package assert

import "testing"

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 0).IsZero()
	assertEmpty(t, buffer)
}

func TestThatIntChainedPrintsNothing(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 0).IsZero().IsEqualTo(0)
	assertEmpty(t, buffer)
}

func TestThatIntChainedPrintsAllMessages(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 2).IsZero().IsEqualTo(1)
	assertContains(t, buffer, "Expected <0>, but was <2>.\nExpected <1>, but was <2>.\n")
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 6).IsZero()
	assertContains(t, buffer, "Expected <0>, but was <6>.\n")
}

func TestThatIntIsEqualTo(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 1).IsEqualTo(1)
	assertEmpty(t, buffer)
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 1).IsEqualTo(2)
	assertContains(t, buffer, "Expected <2>, but was <1>.\n")
}
