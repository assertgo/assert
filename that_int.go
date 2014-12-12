package assert

type ThatInt int

func (thatInt ThatInt) IsZero() {
	if thatInt != 0 {
		panic("that int should be zero")
	}
}
