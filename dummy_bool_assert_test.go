package assert

import "testing"

func TestDummyBoolAssertHasNoBoolAssertRelatedErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(nil).AsBool().
		IsFalse().
		IsTrue()
	mockT.HasErrorMessages(
		"Cannot convert <<nil>> of type <<nil>> to <bool>.",
	)
}
