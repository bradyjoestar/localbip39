package main

import (
	"fmt"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := NewEntropy(256)
	mnemonic, _ := NewMnemonic(entropy)

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
}
