package assert

type Uint64Assert interface {
	IsZero() Uint64Assert
	IsNonZero() Uint64Assert
	// TODO
	// IsEqualTo(expected uint64) Uint64Assert
	// IsNotEqualTo(unexpected uint64) Uint64Assert
	// IsGreaterThan(less uint64) Uint64Assert
	// IsGreaterOrEqualTo(lessOrEqual uint64) Uint64Assert
	// IsLessThan(greater uint64) Uint64Assert
	// IsLessOrEqualTo(greaterOrEqual uint64) Uint64Assert
	// IsBetween(lowerBound, upperBound uint64) Uint64Assert
	// IsNotBetween(lowerBound, upperBound uint64) Uint64Assert
	// IsIn(elements ...uint64) Uint64Assert
	// IsNotIn(elements ...uint64) Uint64Assert
	// IsEven() Uint64Assert
	// IsOdd() Uint64Assert
	// IsDivisibleBy(divisor uint64) Uint64Assert
	// IsNotDivisibleBy(divisor uint64) Uint64Assert
}
