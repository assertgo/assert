package assert

import "testing"

func TestThatIntIsZeroHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsZero().IsZero()
	mockT.HasNoErrors()
}

func TestThatIntIsZeroHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsZero().IsZero()
	mockT.HasErrorMessages(
		"Expected zero, but was <6>.",
		"Expected zero, but was <6>.",
	)
}

func TestThatIntIsNonZeroHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -1).IsNonZero()
	ThatInt(mockT, 1).IsNonZero()
	mockT.HasNoErrors()
}

func TestThatIntIsNonZeroHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonZero().IsNonZero()
	mockT.HasErrorMessages(
		"Expected nonzero, but was zero.",
		"Expected nonzero, but was zero.",
	)
}

func TestThatIntIsEqualToHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsEqualTo(1).IsEqualTo(1)
	mockT.HasNoErrors()
}

func TestThatIntIsEqualToHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsEqualTo(2).IsEqualTo(3)
	mockT.HasErrorMessages(
		"Expected integer equal to <2>, but was <1>.",
		"Expected integer equal to <3>, but was <1>.",
	)
}

func TestThatIntIsNotEqualToHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsNotEqualTo(4).IsNotEqualTo(6)
	mockT.HasNoErrors()
}

func TestThatIntIsNotEqualToHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -5).IsNotEqualTo(-5).IsNotEqualTo(-5)
	mockT.HasErrorMessages(
		"Expected integer not equal to <-5>, but was equal.",
		"Expected integer not equal to <-5>, but was equal.",
	)
}

func TestThatIntIsPositiveHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsPositive()
	ThatInt(mockT, 10).IsPositive()
	mockT.HasNoErrors()
}

func TestThatIntIsPositiveHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsPositive()
	ThatInt(mockT, -10).IsPositive()
	mockT.HasErrorMessages(
		"Expected positive integer, but was <0>.",
		"Expected positive integer, but was <-10>.",
	)
}

func TestThatIntIsNonPositiveHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonPositive()
	ThatInt(mockT, -30).IsNonPositive()
	mockT.HasNoErrors()
}

func TestThatIntIsNonPositiveHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsNonPositive()
	ThatInt(mockT, 30).IsNonPositive()
	mockT.HasErrorMessages(
		"Expected nonpositive integer, but was <1>.",
		"Expected nonpositive integer, but was <30>.",
	)
}

func TestThatIntIsNegativeHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -1).IsNegative()
	ThatInt(mockT, -100).IsNegative()
	mockT.HasNoErrors()
}

func TestThatIntIsNegativeHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNegative()
	ThatInt(mockT, 100).IsNegative()
	mockT.HasErrorMessages(
		"Expected negative integer, but was <0>.",
		"Expected negative integer, but was <100>.",
	)
}

func TestThatIntIsNonNegativeHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNonNegative()
	ThatInt(mockT, 300).IsNonNegative()
	mockT.HasNoErrors()
}

func TestThatIntIsNonNegativeHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -1).IsNonNegative()
	ThatInt(mockT, -300).IsNonNegative()
	mockT.HasErrorMessages(
		"Expected nonnegative integer, but was <-1>.",
		"Expected nonnegative integer, but was <-300>.",
	)
}

func TestThatIntIsGreaterThanHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsGreaterThan(3).IsGreaterThan(-666)
	mockT.HasNoErrors()
}

func TestThatIntIsGreaterThanHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsGreaterThan(4).IsGreaterThan(777)
	mockT.HasErrorMessages(
		"Expected integer greater than <4>, but was <4>.",
		"Expected integer greater than <777>, but was <4>.",
	)
}

func TestThatIntIsGreaterOrEqualToHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -2).IsGreaterOrEqualTo(-2).IsGreaterOrEqualTo(-128)
	mockT.HasNoErrors()
}

func TestThatIntIsGreaterOrEqualToHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -2).IsGreaterOrEqualTo(-1).IsGreaterOrEqualTo(128)
	mockT.HasErrorMessages(
		"Expected integer greater or equal to <-1>, but was <-2>.",
		"Expected integer greater or equal to <128>, but was <-2>.",
	)
}

func TestThatIntIsLessThanHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 28).IsLessThan(29).IsLessThan(999)
	mockT.HasNoErrors()
}

func TestThatIntIsLessThanHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 28).IsLessThan(28).IsLessThan(0)
	mockT.HasErrorMessages(
		"Expected integer lesser than <28>, but was <28>.",
		"Expected integer lesser than <0>, but was <28>.",
	)
}

func TestThatIntIsLessOrEqualToHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -6).IsLessOrEqualTo(-6).IsLessOrEqualTo(6)
	mockT.HasNoErrors()
}

func TestThatIntIsLessOrEqualToHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, -6).IsLessOrEqualTo(-7).IsLessOrEqualTo(-128)
	mockT.HasErrorMessages(
		"Expected integer lesser or equal to <-7>, but was <-6>.",
		"Expected integer lesser or equal to <-128>, but was <-6>.",
	)
}

func TestThatIntIsBetweenHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsBetween(-42, 0).IsBetween(-5, 5).IsBetween(0, 42)
	mockT.HasNoErrors()
}

func TestThatIntIsBetweenHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsBetween(-42, -1).IsBetween(1, 42)
	mockT.HasErrorMessages(
		"Expected integer to be between <-42, -1>, but was <0>.",
		"Expected integer to be between <1, 42>, but was <0>.",
	)
}

func TestThatIntIsNotBetweenHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNotBetween(-42, -1).IsNotBetween(1, 42)
	mockT.HasNoErrors()
}

func TestThatIntIsNotBetweenHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsNotBetween(-42, 0).IsNotBetween(-5, 5).IsNotBetween(0, 42)
	mockT.HasErrorMessages(
		"Expected integer not to be between <-42, 0>, but was <0>.",
		"Expected integer not to be between <-5, 5>, but was <0>.",
		"Expected integer not to be between <0, 42>, but was <0>.",
	)
}

func TestThatIntIsInHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsIn(1, 2, 3, 5, 8, 13).IsIn(2, 3, 5, 7, 11)
	mockT.HasNoErrors()
}

func TestThatIntIsInHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsIn(1, 2, 3, 5, 8, 13).IsIn(2, 3, 5, 7, 11)
	mockT.HasErrorMessages(
		"Expected integer to be in (1, 2, 3, 5, 8, 13), but was <6>.",
		"Expected integer to be in (2, 3, 5, 7, 11), but was <6>.",
	)
}

func TestThatIntIsNotInHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsNotIn(1, 2, 3, 5, 8, 13).IsNotIn(2, 3, 5, 7, 11)
	mockT.HasNoErrors()
}

func TestThatIntIsNotInHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsNotIn(1, 2, 3, 5, 8, 13).IsNotIn(2, 3, 5, 7, 11)
	mockT.HasErrorMessages(
		"Expected integer not to be in (1, 2, 3, 5, 8, 13), but was <5>.",
		"Expected integer not to be in (2, 3, 5, 7, 11), but was <5>.",
	)
}

func TestThatIntIsEvenHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsEven()
	mockT.HasNoErrors()
}

func TestThatIntIsEvenHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsEven().IsEven()
	mockT.HasErrorMessages(
		"Expected integer to be even, but was <5>.",
		"Expected integer to be even, but was <5>.",
	)
}

func TestThatIntIsOddHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 5).IsOdd()
	mockT.HasNoErrors()
}

func TestThatIntIsOddHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsOdd().IsOdd()
	mockT.HasErrorMessages(
		"Expected integer to be odd, but was <6>.",
		"Expected integer to be odd, but was <6>.",
	)
}

func TestThatIntIsDivisibleByHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsDivisibleBy(1).IsDivisibleBy(2)
	mockT.HasNoErrors()
}

func TestThatIntIsDivisibleByHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsDivisibleBy(5).IsDivisibleBy(7)
	mockT.HasErrorMessages(
		"Expected integer to be divisible by <5>, but was <6>.",
		"Expected integer to be divisible by <7>, but was <6>.",
	)
}

func TestThatIntIsNotDivisibleByHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 4).IsNotDivisibleBy(3).IsNotDivisibleBy(5)
	mockT.HasNoErrors()
}

func TestThatIntIsNotDivisibleByHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 6).IsNotDivisibleBy(2).IsNotDivisibleBy(3)
	mockT.HasErrorMessages(
		"Expected integer not to be divisible by <2>, but was <6>.",
		"Expected integer not to be divisible by <3>, but was <6>.",
	)
}
