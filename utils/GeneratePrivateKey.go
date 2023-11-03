package utils

import (
	"crypto/rand"
	"math/big"
)

func GeneratePrivateKey(p *big.Int) (*big.Int, error) {
	one := big.NewInt(1)
	a, err := rand.Int(rand.Reader, p)
	if err != nil {
		return nil, err
	}

	// 确保 a < p-1 且 a 与 p-1 互质
	for a.Cmp(p) >= 0 || new(big.Int).GCD(nil, nil, a, p).Cmp(one) != 0 {
		a, err = rand.Int(rand.Reader, p)
		if err != nil {
			return nil, err
		}
	}
	return a, nil
}
