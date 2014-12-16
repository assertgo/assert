assertgo
========

[![Build Status](https://travis-ci.org/assertgo/assert.svg?branch=develop)](https://travis-ci.org/assertgo/assert) [![Coverage Status](https://coveralls.io/repos/assertgo/assert/badge.png?branch=develop)](https://coveralls.io/r/assertgo/assert?branch=develop)

The coolest assertion library for go language!

```go
func TestPerfectNumber(t *testing.T) {
	assert := assert.Setup(t)
	num := ComputePerfectNumber()
	assert.ThatInt(num).
		IsEqualTo(6).
		IsNonZero().
		IsPositive().
		IsNonNegative().
		IsEven().
		IsDivisibleBy(3).
		IsBetween(4, 10).
		IsIn(496, 28, 6).
		IsNotIn(2, 3, 5, 7, 11, 13)
}
```

Stating what you expect was never that easy.
