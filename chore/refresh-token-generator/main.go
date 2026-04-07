package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	refreshToken := base64.RawURLEncoding.EncodeToString(b)
	fmt.Println("refreshToken: ", refreshToken)

	h := sha256.Sum256([]byte(refreshToken))
	hashedRefreshToken := hex.EncodeToString(h[:])
	fmt.Println("hashedRefreshToken", hashedRefreshToken)

	/*
		refreshToken:  KNmEWXT6vQR4vp2eh2m46cTMR6iPDEl2Qq-K9eNeCyQ
		hashedRefreshToken 1c92ff780fd5bcf3aade663467bcd51bfdc6c85571364f3cb6eac4cc7df9c19a
		// the hashedRefreshToken will always be the same if the refreshToken doesn't change
	*/
}
