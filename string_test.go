package assert

import "testing"

func TestThatStringIsEqualToHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Alice").IsEqualTo("Alice").IsEqualTo("Alice")
	mockT.HasNoErrors()
}

func TestThatStringIsEqualToHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Alice").IsEqualTo("Bob").IsEqualTo("Charlie")
	mockT.HasErrorMessages(
		"Expected <Bob>, but was <Alice>.",
		"Expected <Charlie>, but was <Alice>.",
	)
}

func TestThatStringIsNotEqualToHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Alice").IsNotEqualTo("Bob").IsNotEqualTo("Bob")
	mockT.HasNoErrors()
}

func TestThatStringIsNotEqualToHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Alice").IsNotEqualTo("Alice")
	mockT.HasErrorMessages(
		"Expected string not equal to <Alice>, but was equal.",
	)
}

func TestThatStringIsEmptyHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("").IsEmpty()
	mockT.HasNoErrors()
}

func TestThatStringIsEmptyHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").IsEmpty()
	mockT.HasErrorMessages(
		"Expected string to be empty, but was <Marta>.",
	)
}

func TestThatStringIsNotEmptyHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").IsNotEmpty()
	mockT.HasNoErrors()
}

func TestThatStringIsNotEmptyHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("").IsNotEmpty()
	mockT.HasErrorMessages(
		"Expected string to not be empty, but was.",
	)
}

func TestThatStringIsInSliceHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").IsInSlice([]string{"Marta"})
	mockT.HasNoErrors()
}

func TestThatStringIsInSliceHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").IsInSlice([]string{"Michal", "Macy"})
	mockT.HasErrorMessages(
		"Expected string to be in slice <[Michal Macy]>, but wasn't.",
	)
}

func TestThatStringIsNotInSliceHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").IsNotInSlice([]string{"Michal"})
	mockT.HasNoErrors()
}

func TestThatStringIsNotInSliceHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").IsNotInSlice([]string{"Marta"})
	mockT.HasErrorMessages(
		"Expected string to not be in slice <[Marta]>, but was.",
	)
}

func TestThatStringContainsHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").Contains("art")
	mockT.HasNoErrors()
}

func TestThatStringContainsHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").Contains("Michal")
	mockT.HasErrorMessages(
		"Expected string <Marta> to contain <Michal>, but didn't.",
	)
}

func TestThatStringDoesNotContainHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").DoesNotContain("Michal")
	mockT.HasNoErrors()
}

func TestThatStringDoesNotContainHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").DoesNotContain("art")
	mockT.HasErrorMessages(
		"Expected string <Marta> to not contain <art>, but it did.",
	)
}

func TestThatStringIsLowerCaseHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("marta").IsLowerCase()
	mockT.HasNoErrors()
}

func TestThatStringIsLowerCaseHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Marta").IsLowerCase()
	mockT.HasErrorMessages(
		"Expected string <Marta> to be lower case, but wasn't.",
	)
}

func TestThatStringIsUpperCaseHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("MARTA").IsUpperCase()
	mockT.HasNoErrors()
}

func TestThatStringIsUpperCaseHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("MArTa").IsUpperCase()
	mockT.HasErrorMessages(
		"Expected string <MArTa> to be upper case, but wasn't.",
	)
}

func TestThatStringIsNotLowerCaseHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("MArta").IsNotLowerCase()
	mockT.HasNoErrors()
}

func TestThatStringIsNotLowerCaseHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("marta").IsNotLowerCase()
	mockT.HasErrorMessages(
		"Expected string <marta> to not be lower case, but was.",
	)
}

func TestThatStringIsNotUpperCaseHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("MArta").IsNotUpperCase()
	mockT.HasNoErrors()
}

func TestThatStringIsNotUpperCaseHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("MARTA").IsNotUpperCase()
	mockT.HasErrorMessages(
		"Expected string <MARTA> to not be upper case, but was.",
	)
}

func TestThatStringMatchesHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Ewelina").Matches("^Ewe(\\w+)")
	mockT.HasNoErrors()
}

func TestThatStringMatchesHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatString("Ewelina").Matches("ewelin?")
	mockT.HasErrorMessages(
		"Expected string <Ewelina> to match <ewelin?>, but didn't.",
	)
}
