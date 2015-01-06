assertgo ![](https://raw.githubusercontent.com/assertgo/icon/master/assertgo_64.png)
====================================================================================

[![Build Status](https://travis-ci.org/assertgo/assert.svg?branch=develop)](https://travis-ci.org/assertgo/assert)
[![Coverage Status](https://img.shields.io/coveralls/assertgo/assert.svg)](https://coveralls.io/r/assertgo/assert)
[![GoDoc](https://godoc.org/github.com/assertgo/assert?status.svg)](https://godoc.org/github.com/assertgo/assert)

The coolest assertion library for go language!

Version: 2.0.0

Getting started
---------------

To get the package, execute:

	go get github.com/assertgo/assert

To import this package, add the following line to your code:

	import "github.com/assertgo/assert"

And just start using it.

```go
func TestPerfectNumber(t *testing.T) {
	assert := assert.New(t)
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

Contibuting
-----------

Please see [CONTRIBUTING.md](CONTRIBUTING.md).

License
-------

assertgo is a free software and may be used under the terms specified in the [LICENSE](LICENSE) file.
