package assert

import "testing"

func TestThatStringIsEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsEqualTo("Alice").IsEqualTo("Alice")
	mockT.HasNoErrors()
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsEqualTo("Bob").IsEqualTo("Charlie")
	mockT.HasErrorMessages(
		"Expected <Bob>, but was <Alice>.",
		"Expected <Charlie>, but was <Alice>.",
	)
}

func TestThatStringIsNotEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsNotEqualTo("Bob").IsNotEqualTo("Bob")
	mockT.HasNoErrors()
}

func TestThatStringIsNotEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsNotEqualTo("Alice")
	mockT.HasErrorMessages(
		"Expected string not equal to <Alice>, but was equal.",
	)
}

func TestThatStringIsEmptyPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "").IsEmpty()
	mockT.HasNoErrors()
}

func TestThatStringIsEmptyPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsEmpty()
	mockT.HasErrorMessages(
		"Expected string to be empty, but was <Marta>.",
	)
}

func TestThatStringIsInSlicePrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsInSlice([]string{"Marta"})
	mockT.HasNoErrors()
}

func TestThatStringIsInSlicePrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsInSlice([]string{"Michal", "Macy"})
	mockT.HasErrorMessages(
		"Expected string to be in slice <[Michal Macy]>, but wasn't.",
	)
}

func TestThatStringContainsePrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").Contains("art")
	mockT.HasNoErrors()
}

func TestThatStringContainsPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").Contains("Michal")
	mockT.HasErrorMessages(
		"Expected string to contain <Michal>, but didn't.",
	)
}
