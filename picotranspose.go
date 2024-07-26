// The pattern to crack this is
// every 3rd char needs to be shifted 2 places backwards
// so for example: heT -> The
// both h and e shifted up 1 index and T moved -2
// we iterate through the string and do this for every 3rd char

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func moveElement(slice []string, fromIndex, toIndex int) string {
	// create a new private storage slice for this block
	resultSlice := make([]string, len(slice))
	// shift the values
	resultSlice[toIndex] = slice[fromIndex]
	copy(resultSlice[toIndex+1:], slice[toIndex:fromIndex])

	// send a string object instead of a slice
	return strings.Join(resultSlice, "")
}

func main() {
	// read from challenge file into a []byte var
	cipherByte, err := os.ReadFile("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Convert file string to slice for modifications
	cipherSlice := strings.Split(string(cipherByte), "")

	// create a slice that we can append each block to
	decodeSlice := make([]string, len(cipherSlice))

	// loop through string and shift indexes with moveElement
	for i := 2; i <= len(cipherSlice); i += 3 {
		decodeSlice = append(decodeSlice, moveElement(cipherSlice, i, i-2), "")
	}

	// assemble the slice into a string and print
	fmt.Println(strings.Join(decodeSlice, ""))

}
