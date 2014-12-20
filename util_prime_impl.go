package assert

func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	v64 := int64(value)
	for i := int64(5); i*i <= v64; i += 6 {
		if v64%i == 0 || v64%(i+2) == 0 {
			return false
		}
	}
	return true
}
