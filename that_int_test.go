package assert

import "testing"

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsZero().IsZero()
	mockT.HasNoErrors()
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsZero()
	mockT.HasErrorMessages(
		"Expected zero, but was <6>.")
}

func TestThatIntIsNonzeroPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -1).IsNonzero()
	ThatInt(mockT, 1).IsNonzero()
	mockT.HasNoErrors()
}

func TestThatIntIsNonzeroPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonzero().IsNonzero()
	mockT.HasErrorMessages(
		"Expected nonzero, but was zero.",
		"Expected nonzero, but was zero.")
}

func TestThatIntIsEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsEqualTo(1).IsEqualTo(1)
	mockT.HasNoErrors()
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsEqualTo(2)
	mockT.HasErrorMessages(
		"Expected integer equal to <2>, but was <1>.")
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
	mockT.HasErrorMessages(
		"Expected positive integer, but was <0>.",
		"Expected positive integer, but was <-10>.")
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
	mockT.HasErrorMessages(
		"Expected negative integer, but was <0>.",
		"Expected negative integer, but was <100>.")
}

func TestThatIntIsGreaterThanPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsGreaterThan(3).IsGreaterThan(-666)
	mockT.HasNoErrors()
}

func TestThatIntIsGreaterThanPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsGreaterThan(4).IsGreaterThan(777)
	mockT.HasErrorMessages(
		"Expected integer greater than <4>, but was <4>.",
		"Expected integer greater than <777>, but was <4>.")
}

func TestThatIntIsGreaterOrEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -2).IsGreaterOrEqualTo(-2).IsGreaterOrEqualTo(-128)
	mockT.HasNoErrors()
}

func TestThatIntIsGreaterOrEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -2).IsGreaterOrEqualTo(-1).IsGreaterOrEqualTo(128)
	mockT.HasErrorMessages(
		"Expected integer greater or equal to <-1>, but was <-2>.",
		"Expected integer greater or equal to <128>, but was <-2>.")
}
