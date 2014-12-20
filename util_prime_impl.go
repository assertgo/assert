package assert

import (
	"io/ioutil"
	"log"
	"math/big"

	"github.com/akalin/aks-go/aks"
)

func isPrime(value int) bool {
	if value < 2 {
		return false
	}
	n := big.NewInt(int64(value))
	jobs := 1
	r := aks.CalculateAKSModulus(n)
	M := aks.CalculateAKSUpperBound(n, r)
	logger := log.New(ioutil.Discard, "", 0)
	a := aks.GetAKSWitness(n, r, &big.Int{}, M, jobs, logger)
	return a == nil
}
