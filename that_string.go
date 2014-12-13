package assert

type ThatString string

func (thatString ThatString) IsEqualTo(comparedString string) {
	if string(thatString) != comparedString {
		panic("")
	}
}
