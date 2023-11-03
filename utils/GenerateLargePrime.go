package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateLargePrime
// param:
// bit: number of prime
func GenerateLargePrime(bits int) (*big.Int, error) {

	prime, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return prime, nil
}
