package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible bas64.
	// Here's how to encode using the standard encoder. The encoder requires a []byte so we convert our string
	// to that type
	senc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(senc)

	// Decoding may return an error, which you can check for if you don't already if the input is well-formed.
	sdec, _ := b64.StdEncoding.DecodeString(senc)
	fmt.Println(string(sdec))
	fmt.Println()

	// This example encodes/decodes using URL-compatible base64 format.
	uenc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(string(uenc))
	udec, _ := b64.URLEncoding.DecodeString(senc)
	fmt.Println(string(udec))

	// The string encodes to slightly different values with the standard- and URL base64 encoders
	// (training + vs. -) but they both decode to the original string as desired.
}
