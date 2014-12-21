package assert

import "testing"

func TestThatIsEqualToHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.That(TestThatIsEqualToHasNoErrors).
		IsEqualTo(TestThatIsEqualToHasNoErrors).
		IsEqualTo(TestThatIsEqualToHasNoErrors)
	mockT.HasNoErrors()
}
