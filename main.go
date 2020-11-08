package main

import (
	"fmt"
	//"bytes";
	"encoding/hex"
)

func main() {
	//h := []byte("4927")
	fmt.Println(hex.DecodeString("4927"))
	fmt.Println([]byte("I'"))
	// omg I'm decoding the characters in the string to bytes
}
