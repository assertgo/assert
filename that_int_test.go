package assert

import "testing"

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{}
	ThatInt(mockT, 0).IsZero()
	assertEmpty(t, mockT)
}

func TestThatIntChainedPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{}
	ThatInt(mockT, 0).IsZero().IsEqualTo(0)
	assertEmpty(t, mockT)
}

func TestThatIntChainedPrintsAllMessages(t *testing.T) {
	mockT := &mockTestingT{}
	ThatInt(mockT, 2).IsZero().IsEqualTo(1)
	assertContains(t, mockT, "Expected <0>, but was <2>.\nExpected <1>, but was <2>.\n")
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	mockT := &mockTestingT{}
	ThatInt(mockT, 6).IsZero()
	assertContains(t, mockT, "Expected <0>, but was <6>.\n")
}

func TestThatIntIsEqualToPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{}
	ThatInt(mockT, 1).IsEqualTo(1)
	assertEmpty(t, mockT)
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	mockT := &mockTestingT{}
	ThatInt(mockT, 1).IsEqualTo(2)
	assertContains(t, mockT, "Expected <2>, but was <1>.\n")
}
