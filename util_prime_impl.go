package assert

func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	uns := uint(value)
	for i := uint(5); i*i <= uns; i += 6 {
		if uns%i == 0 || uns%(i+2) == 0 {
			return false
		}
	}
	return true
}
