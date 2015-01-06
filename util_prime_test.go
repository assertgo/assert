package assert

import "testing"

func TestLessThanTwoIsNotPrime(t *testing.T) {
	assert := New(t)
	assert.ThatBool(isPrime(-1000)).IsFalse()
	assert.ThatBool(isPrime(-100)).IsFalse()
	assert.ThatBool(isPrime(-10)).IsFalse()
	assert.ThatBool(isPrime(-1)).IsFalse()
	assert.ThatBool(isPrime(0)).IsFalse()
	assert.ThatBool(isPrime(1)).IsFalse()
}

func TestSmallNumbersArePrime(t *testing.T) {
	assert := New(t)
	assert.ThatBool(isPrime(2)).IsTrue()
	assert.ThatBool(isPrime(3)).IsTrue()
	assert.ThatBool(isPrime(5)).IsTrue()
	assert.ThatBool(isPrime(7)).IsTrue()
}

func TestSmallNumbersAreNotPrime(t *testing.T) {
	assert := New(t)
	assert.ThatBool(isPrime(4)).IsFalse()
	assert.ThatBool(isPrime(6)).IsFalse()
	assert.ThatBool(isPrime(8)).IsFalse()
	assert.ThatBool(isPrime(9)).IsFalse()
}

func TestBigNumbersArePrime(t *testing.T) {
	assert := New(t)
	assert.ThatBool(isPrime(2311)).IsTrue()
	assert.ThatBool(isPrime(3203)).IsTrue()
	assert.ThatBool(isPrime(1321)).IsTrue()
}

func TestBigNumbersAreNotPrime(t *testing.T) {
	assert := New(t)
	assert.ThatBool(isPrime(1001)).IsFalse()
	assert.ThatBool(isPrime(5251)).IsFalse()
	assert.ThatBool(isPrime(2809)).IsFalse()
}

func TestReallyBigNumberIsPrime(t *testing.T) {
	assert := New(t)
	assert.ThatBool(isPrime(2147483647)).IsTrue()
}

func TestReallyBigNumberIsNotPrime(t *testing.T) {
	assert := New(t)
	assert.ThatBool(isPrime(46337 * 46337)).IsFalse()
}
