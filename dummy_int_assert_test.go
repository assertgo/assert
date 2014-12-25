package assert

import "testing"

func TestDummyIntAssertHasNoIntAssertRelatedErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(nil).AsInt().
		IsZero().
		IsNonZero().
		IsEqualTo(0).
		IsNotEqualTo(0).
		IsPositive().
		IsNonPositive().
		IsNegative().
		IsNonNegative().
		IsGreaterThan(0).
		IsGreaterOrEqualTo(0).
		IsLessThan(0).
		IsLessOrEqualTo(0).
		IsBetween(0, 1).
		IsNotBetween(0, 1).
		IsIn(0, 1).
		IsNotIn(0, 1).
		IsEven().
		IsOdd().
		IsDivisibleBy(3).
		IsNotDivisibleBy(3).
		IsPrime().
		IsNotPrime().
		IsAnswerToTheUltimateQuestionOfLife()
	mockT.HasErrorMessages(
		"Cannot convert <<nil>> of type <<nil>> to <int>.",
	)
}
