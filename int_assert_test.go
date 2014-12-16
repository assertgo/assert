package assert

import "testing"

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsZero().IsZero()
	mockT.HasNoErrors()
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsZero().IsZero()
	mockT.HasErrorMessages(
		"Expected zero, but was <6>.",
		"Expected zero, but was <6>.",
	)
}

func TestThatIntIsNonZeroPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -1).IsNonZero()
	ThatInt(mockT, 1).IsNonZero()
	mockT.HasNoErrors()
}

func TestThatIntIsNonZeroPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonZero().IsNonZero()
	mockT.HasErrorMessages(
		"Expected nonzero, but was zero.",
		"Expected nonzero, but was zero.",
	)
}

func TestThatIntIsEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsEqualTo(1).IsEqualTo(1)
	mockT.HasNoErrors()
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsEqualTo(2).IsEqualTo(3)
	mockT.HasErrorMessages(
		"Expected integer equal to <2>, but was <1>.",
		"Expected integer equal to <3>, but was <1>.",
	)
}

func TestThatIntIsNotEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsNotEqualTo(4).IsNotEqualTo(6)
	mockT.HasNoErrors()
}

func TestThatIntIsNotEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -5).IsNotEqualTo(-5).IsNotEqualTo(-5)
	mockT.HasErrorMessages(
		"Expected integer not equal to <-5>, but was equal.",
		"Expected integer not equal to <-5>, but was equal.",
	)
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
		"Expected positive integer, but was <-10>.",
	)
}

func TestThatIntIsNonPositivePrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonPositive()
	ThatInt(mockT, -30).IsNonPositive()
	mockT.HasNoErrors()
}

func TestThatIntIsNonPositivePrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsNonPositive()
	ThatInt(mockT, 30).IsNonPositive()
	mockT.HasErrorMessages(
		"Expected nonpositive integer, but was <1>.",
		"Expected nonpositive integer, but was <30>.",
	)
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
		"Expected negative integer, but was <100>.",
	)
}

func TestThatIntIsNonNegativePrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonNegative()
	ThatInt(mockT, 300).IsNonNegative()
	mockT.HasNoErrors()
}

func TestThatIntIsNonNegativePrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -1).IsNonNegative()
	ThatInt(mockT, -300).IsNonNegative()
	mockT.HasErrorMessages(
		"Expected nonnegative integer, but was <-1>.",
		"Expected nonnegative integer, but was <-300>.",
	)
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
		"Expected integer greater than <777>, but was <4>.",
	)
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
		"Expected integer greater or equal to <128>, but was <-2>.",
	)
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
		"Expected integer lesser than <0>, but was <28>.",
	)
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
		"Expected integer lesser or equal to <-128>, but was <-6>.",
	)
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
		"Expected integer to be between <1, 42>, but was <0>.",
	)
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
		"Expected integer not to be between <0, 42>, but was <0>.",
	)
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
		"Expected integer to be in (2, 3, 5, 7, 11), but was <6>.",
	)
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
		"Expected integer not to be in (2, 3, 5, 7, 11), but was <5>.",
	)
}

func TestThatIntIsEvenPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsEven()
	mockT.HasNoErrors()
}

func TestThatIntIsEvenPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsEven().IsEven()
	mockT.HasErrorMessages(
		"Expected integer to be even, but was <5>.",
		"Expected integer to be even, but was <5>.",
	)
}

func TestThatIntIsOddPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsOdd()
	mockT.HasNoErrors()
}

func TestThatIntIsOddPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsOdd().IsOdd()
	mockT.HasErrorMessages(
		"Expected integer to be odd, but was <6>.",
		"Expected integer to be odd, but was <6>.",
	)
}

func TestThatIntIsDivisibleByPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsDivisibleBy(1).IsDivisibleBy(2)
	mockT.HasNoErrors()
}

func TestThatIntIsDivisibleByPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsDivisibleBy(5).IsDivisibleBy(7)
	mockT.HasErrorMessages(
		"Expected integer to be divisible by <5>, but was <6>.",
		"Expected integer to be divisible by <7>, but was <6>.",
	)
}

func TestThatIntIsNotDivisibleByPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsNotDivisibleBy(3).IsNotDivisibleBy(5)
	mockT.HasNoErrors()
}

func TestThatIntIsNotDivisibleByPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsNotDivisibleBy(2).IsNotDivisibleBy(3)
	mockT.HasErrorMessages(
		"Expected integer not to be divisible by <2>, but was <6>.",
		"Expected integer not to be divisible by <3>, but was <6>.",
	)
}
