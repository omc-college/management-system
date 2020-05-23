package handlers

import (
	"math/big"
	"crypto/rand"
	"strconv"
)

// Slice symbols contains symbols which meets in the random string
var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomGenerate returns random string with specified size
func RandomGenerate(n int) string {
	b := make([]rune, n)

	for i := range b {
		length := int64(len(symbols))
		a, _ := rand.Int(rand.Reader,big.NewInt(length))
		c, _ := strconv.ParseInt(a.String(), 10, 0)
		b[i] = symbols[c]
	}

	return string(b)
}