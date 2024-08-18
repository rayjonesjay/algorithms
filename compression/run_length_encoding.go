package compression

import (
	"strconv"
	"strings"
)

/*
Lossless Data Compression - is a technique where data is compressed in such a way the original data can be perfectly constructed
from the compressed data.

Run-Length Encoding(RLE) - is a simple form of lossless data compression where consucutive occurrences
of the same data value (a run) are stored as a single data value and a count(length of run).
*/
func RunLengthEncode(input string) string {
	// construct our encoded data using a builder
	var encoded strings.Builder

	// get length of input
	length := len(input)

	// traverse the input
	for i := 0; i < length; i++ {
		count := 1
		// count the number of times each character repeats
		for i < length-1 && input[i] == input[i+1] {
			count++
			i++
		}
		// write the run length and character
		encoded.WriteString(strconv.Itoa(count))
		encoded.WriteByte(input[i])
	}
	return encoded.String()
}

// RunLengthDecode decodes 'encoded' and returns 'decoded'
func RunLengthDecode(encoded string) string {
	var decoded strings.Builder

	length := len(encoded)
	i := 0
	for i < length {
		numberStartCount := i
		for i < length && encoded[i] >= '0' && encoded[i] <= '9' {
			i++
		}
		count, _ := strconv.Atoi(string(encoded[numberStartCount:i]))

		if i < length {
			currentChar := encoded[i]
			for j := 0; j < count; j++ {
				decoded.WriteByte(currentChar)
			}
		}
		i++
	}
	return decoded.String()
}
