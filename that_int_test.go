package assert

import "testing"

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{t: t}
	ThatInt(mockT, 0).IsZero()
	mockT.assertEmpty()
}

func TestThatIntChainedPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{t: t}
	ThatInt(mockT, 0).IsZero().IsEqualTo(0)
	mockT.assertEmpty()
}

func TestThatIntChainedPrintsAllMessages(t *testing.T) {
	mockT := &mockTestingT{t: t}
	ThatInt(mockT, 2).IsZero().IsEqualTo(1)
	mockT.assertContains("Expected <0>, but was <2>.\nExpected <1>, but was <2>.\n")
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	mockT := &mockTestingT{t: t}
	ThatInt(mockT, 6).IsZero()
	mockT.assertContains("Expected <0>, but was <6>.\n")
}

func TestThatIntIsEqualToPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{t: t}
	ThatInt(mockT, 1).IsEqualTo(1)
	mockT.assertEmpty()
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	mockT := &mockTestingT{t: t}
	ThatInt(mockT, 1).IsEqualTo(2)
	mockT.assertContains("Expected <2>, but was <1>.\n")
}
