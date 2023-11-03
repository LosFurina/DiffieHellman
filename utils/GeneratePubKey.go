package utils

import "math/big"

func GeneratePubKey(g, p, a *big.Int) *big.Int {
	A := new(big.Int).Exp(g, a, p)
	return A
}
