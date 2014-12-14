package assert

import "testing"

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsZero()
	mockT.HasNoErrors()
}

func TestThatIntChainedPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsZero().IsEqualTo(0)
	mockT.HasNoErrors()
}

func TestThatIntChainedPrintsAllMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 2).IsZero().IsEqualTo(1)
	mockT.HasErrorMessage("Expected <0>, but was <2>.\nExpected <1>, but was <2>.\n")
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsZero()
	mockT.HasErrorMessage("Expected <0>, but was <6>.\n")
}

func TestThatIntIsEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsEqualTo(1)
	mockT.HasNoErrors()
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsEqualTo(2)
	mockT.HasErrorMessage("Expected <2>, but was <1>.\n")
}

func TestThatIntIsPositivePrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsPositive()
	ThatInt(mockT, 10).IsPositive()
	mockT.HasNoErrors()
}

func TestThatIntIsPositivePrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsPositive()
	ThatInt(mockT, -10).IsPositive()
	mockT.HasErrorMessage("Expected positive integer, but was <0>.\nExpected positive integer, but was <-10>.\n")
}

func TestThatIntIsNegativePrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -1).IsNegative()
	ThatInt(mockT, -100).IsNegative()
	mockT.HasNoErrors()
}

func TestThatIntIsNegativePrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNegative()
	ThatInt(mockT, 100).IsNegative()
	mockT.HasErrorMessage("Expected negative integer, but was <0>.\nExpected negative integer, but was <100>.\n")
}

func TestThatIntIsGreaterThanPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsGreaterThan(3).IsGreaterThan(-666)
	mockT.HasNoErrors()
}

func TestThatIntIsGreaterThanPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsGreaterThan(4).IsGreaterThan(777)
	mockT.HasErrorMessage("Expected integer greater than <4>, but was <4>.\nExpected integer greater than <777>, but was <4>.\n")
}

func TestThatIntIsGreaterOrEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -2).IsGreaterOrEqualTo(-2).IsGreaterOrEqualTo(-128)
	mockT.HasNoErrors()
}

func TestThatIntIsGreaterOrEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -2).IsGreaterOrEqualTo(-1).IsGreaterOrEqualTo(128)
	mockT.HasErrorMessage("Expected integer greater or equal to <-1>, but was <-2>.\nExpected integer greater or equal to <128>, but was <-2>.\n")
}
