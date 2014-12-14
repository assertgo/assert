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

func TestThatIntIsNotEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsNotEqualTo(4).IsNotEqualTo(6)
	mockT.HasNoErrors()
}

func TestThatIntIsNotEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -5).IsNotEqualTo(-5)
	mockT.HasErrorMessages(
		"Expected integer not equal to <-5>, but was equal.")
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

func TestThatIntIsNonpositivePrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonpositive()
	ThatInt(mockT, -30).IsNonpositive()
	mockT.HasNoErrors()
}

func TestThatIntIsNonpositivePrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsNonpositive()
	ThatInt(mockT, 30).IsNonpositive()
	mockT.HasErrorMessages(
		"Expected nonpositive integer, but was <1>.",
		"Expected nonpositive integer, but was <30>.")
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

func TestThatIntIsNonnegativePrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonnegative()
	ThatInt(mockT, 300).IsNonnegative()
	mockT.HasNoErrors()
}

func TestThatIntIsNonnegativePrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -1).IsNonnegative()
	ThatInt(mockT, -300).IsNonnegative()
	mockT.HasErrorMessages(
		"Expected nonnegative integer, but was <-1>.",
		"Expected nonnegative integer, but was <-300>.")
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

func TestThatIntIsLessThanPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 28).IsLessThan(29).IsLessThan(999)
	mockT.HasNoErrors()
}

func TestThatIntIsLessThanPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 28).IsLessThan(28).IsLessThan(0)
	mockT.HasErrorMessages(
		"Expected integer lesser than <28>, but was <28>.",
		"Expected integer lesser than <0>, but was <28>.")
}

func TestThatIntIsLessOrEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -6).IsLessOrEqualTo(-6).IsLessOrEqualTo(6)
	mockT.HasNoErrors()
}

func TestThatIntIsLessOrEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -6).IsLessOrEqualTo(-7).IsLessOrEqualTo(-128)
	mockT.HasErrorMessages(
		"Expected integer lesser or equal to <-7>, but was <-6>.",
		"Expected integer lesser or equal to <-128>, but was <-6>.")
}

func TestThatIntIsBetweenPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsBetween(-42, 0).IsBetween(-5, 5).IsBetween(0, 42)
	mockT.HasNoErrors()
}

func TestThatIntIsBetweenPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsBetween(-42, -1).IsBetween(1, 42)
	mockT.HasErrorMessages(
		"Expected integer to be between <-42, -1>, but was <0>.",
		"Expected integer to be between <1, 42>, but was <0>.")
}

func TestThatIntIsNotBetweenPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNotBetween(-42, -1).IsNotBetween(1, 42)
	mockT.HasNoErrors()
}

func TestThatIntIsNotBetweenPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNotBetween(-42, 0).IsNotBetween(-5, 5).IsNotBetween(0, 42)
	mockT.HasErrorMessages(
		"Expected integer not to be between <-42, 0>, but was <0>.",
		"Expected integer not to be between <-5, 5>, but was <0>.",
		"Expected integer not to be between <0, 42>, but was <0>.")
}

func TestThatIntIsInPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsIn(1, 2, 3, 5, 8, 13).IsIn(2, 3, 5, 7, 11)
	mockT.HasNoErrors()
}

func TestThatIntIsInPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsIn(1, 2, 3, 5, 8, 13).IsIn(2, 3, 5, 7, 11)
	mockT.HasErrorMessages(
		"Expected integer to be in (1, 2, 3, 5, 8, 13), but was <6>.",
		"Expected integer to be in (2, 3, 5, 7, 11), but was <6>.")
}

func TestThatIntIsNotInPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsNotIn(1, 2, 3, 5, 8, 13).IsNotIn(2, 3, 5, 7, 11)
	mockT.HasNoErrors()
}

func TestThatIntIsNotInPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsNotIn(1, 2, 3, 5, 8, 13).IsNotIn(2, 3, 5, 7, 11)
	mockT.HasErrorMessages(
		"Expected integer not to be in (1, 2, 3, 5, 8, 13), but was <5>.",
		"Expected integer not to be in (2, 3, 5, 7, 11), but was <5>.")
}
