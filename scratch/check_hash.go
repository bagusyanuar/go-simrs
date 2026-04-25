package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	v := "some_random_verifier_string"
	h := sha256.Sum256([]byte(v))
	s := base64.RawURLEncoding.EncodeToString(h[:])
	fmt.Printf("Verifier: %s\n", v)
	fmt.Printf("Challenge: %s\n", s)
}
