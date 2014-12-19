package assert

import "testing"

func TestThatStringIsEqualToHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsEqualTo("Alice").IsEqualTo("Alice")
	mockT.HasNoErrors()
}

func TestThatStringIsEqualToHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsEqualTo("Bob").IsEqualTo("Charlie")
	mockT.HasErrorMessages(
		"Expected <Bob>, but was <Alice>.",
		"Expected <Charlie>, but was <Alice>.",
	)
}

func TestThatStringIsNotEqualToHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsNotEqualTo("Bob").IsNotEqualTo("Bob")
	mockT.HasNoErrors()
}

func TestThatStringIsNotEqualToHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsNotEqualTo("Alice")
	mockT.HasErrorMessages(
		"Expected string not equal to <Alice>, but was equal.",
	)
}

func TestThatStringIsEmptyHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "").IsEmpty()
	mockT.HasNoErrors()
}

func TestThatStringIsEmptyHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsEmpty()
	mockT.HasErrorMessages(
		"Expected string to be empty, but was <Marta>.",
	)
}

func TestThatStringIsInSliceHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsInSlice([]string{"Marta"})
	mockT.HasNoErrors()
}

func TestThatStringIsInSliceHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsInSlice([]string{"Michal", "Macy"})
	mockT.HasErrorMessages(
		"Expected string to be in slice <[Michal Macy]>, but wasn't.",
	)
}

func TestThatStringIsNotInSliceHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsNotInSlice([]string{"Michal"})
	mockT.HasNoErrors()
}

func TestThatStringIsNotInSliceHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsNotInSlice([]string{"Marta"})
	mockT.HasErrorMessages(
		"Expected string to not be in slice <[Marta]>, but was.",
	)
}

func TestThatStringContainsHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").Contains("art")
	mockT.HasNoErrors()
}

func TestThatStringContainsHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").Contains("Michal")
	mockT.HasErrorMessages(
		"Expected string to contain <Michal>, but didn't.",
	)
}
