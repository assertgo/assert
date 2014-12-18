package assert

func isPrime(value int) bool {
	if value < 2 {
		return false
	}
	return value < 3
}
