package main

import (
	"fmt"
	"github.com/LosFurina/DiffieHellman/utils"
)

func main() {
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

	b, err := utils.GeneratePrivateKey(p)
	if err != nil {
		return
	}
	B := utils.GeneratePubKey(g, p, b)
	fmt.Println("B is:", B)

	SharedKeyA := utils.GenerateSharedKey(A, p, b)
	SharedKeyB := utils.GenerateSharedKey(B, p, a)

	fmt.Println("Shared Key A:", SharedKeyA)
	fmt.Println("Shared Key B:", SharedKeyB)

	if SharedKeyA.Cmp(SharedKeyB) == 0 {
		fmt.Println("Shared keys match. Secure connection established.")
	} else {
		fmt.Println("Shared keys do not match. Connection compromised.")
	}
}
