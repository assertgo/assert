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

func TestThatStringIsNotEmptyHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsNotEmpty()
	mockT.HasNoErrors()
}

func TestThatStringIsNotEmptyHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "").IsNotEmpty()
	mockT.HasErrorMessages(
		"Expected string to not be empty, but was.",
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
		"Expected string <Marta> to contain <Michal>, but didn't.",
	)
}

func TestThatStringDoesNotContainHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").DoesNotContain("Michal")
	mockT.HasNoErrors()
}

func TestThatStringDoesNotContainHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").DoesNotContain("art")
	mockT.HasErrorMessages(
		"Expected string <Marta> to not contain <art>, but it did.",
	)
}

func TestThatStringIsLowerCaseHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "marta").IsLowerCase()
	mockT.HasNoErrors()
}

func TestThatStringIsLowerCaseHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Marta").IsLowerCase()
	mockT.HasErrorMessages(
		"Expected string <Marta> to be lower case, but wasn't.",
	)
}
