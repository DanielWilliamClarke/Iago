package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/DanielWilliamClarke/Iago/lib"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func main() {
	n := 1024
	b, err := GenerateRandomBytes(n)
	if err != nil {
		log.Fatalf("error encountered %v", err)
	}
	str := base64.URLEncoding.EncodeToString(b)
	substring := lib.ExtractLargestUniqueSubstring(str)
	fmt.Printf(
		"%s\n\ncontains a unique substring %s of length %d",
		str, substring, len(substring))
}
