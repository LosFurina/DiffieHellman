package test

import (
	"fmt"
	"github.com/LosFurina/DiffieHellman/utils"
	"testing"
)

func TestFindPrimitiveRoot(t *testing.T) {
	prime, err := utils.GenerateLargePrime(2048)
	fmt.Println("p is:", prime)
	if err != nil {
		return
	}
	res := utils.FindPrimitiveRoot(prime)
	fmt.Println("g is:", res)
}

func TestGeneratePrivateKey(t *testing.T) {
	P, err := utils.GenerateLargePrime(2048)
	if err != nil {
		return
	}
	a, err := utils.GeneratePrivateKey(P)
	if err != nil {
		return
	}
	fmt.Println("a is:", a)
}

func TestGeneratePubKey(t *testing.T) {
	p, err := utils.GenerateLargePrime(2048)
	if err != nil {
		return
	}
	g := utils.FindPrimitiveRoot(p)

	a, err := utils.GeneratePrivateKey(p)
	if err != nil {
		return
	}
	A := utils.GeneratePubKey(g, p, a)
	fmt.Println("A is:", A)
}
