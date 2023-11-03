package utils

import "math/big"

func GenerateSharedKey(PubKey, p, PriKey *big.Int) *big.Int {
	sharedKey := new(big.Int).Exp(PubKey, PriKey, p)
	return sharedKey
}
