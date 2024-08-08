package main

import (
	"fmt"

	"algo/compression"
)

func main() {
	data := "AABBBCCCDDDDDDDDDDDDDDD"
	fmt.Println("original", data)
	encoded := compression.RunLengthEncode(data)
	fmt.Println("encoded:", encoded)
	decoded := compression.RunLengthDecode(encoded)
	fmt.Println("decoded:", decoded)
}
