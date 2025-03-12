package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

// func GenerateUrl(urlLength int) string {
// 	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	result := ""

// 	for i := 0; i < urlLength; i++ {
// 		result += string(characters[rand.Intn(len(characters)-1)])
// 	}
// 	return result
// }

func GenerateUrl2(urlLength int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, urlLength)

	for i := 0; i < urlLength; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		result[i] = characters[n.Int64()]
	}

	return string(result)
}

func ShortenURLHash(url string) string {
	hash := sha256.Sum256([]byte(url))
	shortKey := hex.EncodeToString(hash[:4]) // Take only the first 8 characters (4 bytes)
	return shortKey
}
