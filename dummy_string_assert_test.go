package assert

import "testing"

func TestDummyStringAssertHasNoStringAssertRelatedErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(nil).AsString().
		IsEqualTo("").
		IsNotEqualTo("").
		IsEmpty().
		IsNotEmpty().
		IsInSlice([]string{"", "abc"}).
		IsNotInSlice([]string{"", "abc"}).
		Contains("a").
		DoesNotContain("a").
		IsLowerCase().
		IsNotLowerCase().
		IsUpperCase().
		IsNotUpperCase()
	mockT.HasErrorMessages(
		"Cannot convert <<nil>> of type <<nil>> to <string>.",
	)
}
