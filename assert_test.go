package assert

import "testing"

func TestCanCreateAssert(t *testing.T) {
	mockT := newMockT(t)
	assert := Setup(mockT)
	assert.ThatInt(0).IsZero()
	assert.ThatString("").IsEmpty()
	mockT.HasNoErrors()
}
