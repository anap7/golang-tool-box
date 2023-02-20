package main

import (
	"fmt"
	"intro-test/address"
)

func main() {
	addressType := address.AddressType("Rua show!")
	fmt.Println(addressType)
}