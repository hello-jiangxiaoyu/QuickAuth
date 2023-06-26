package safe

import (
	"crypto/rand"
)

const Base64Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
const Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const HexChars = "0123456789abcdef"
const DecChars = "0123456789"

func Rand64(n int) string  { return RandString(n, Base64Chars) }
func Rand62(n int) string  { return RandString(n, Base62Chars) }
func RandDec(n int) string { return RandString(n, DecChars) }
func RandHex(n int) string { return RandString(n, HexChars) }

// RandString Generate random string
func RandString(n int, letters string) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	if len(letters) == 0 {
		letters = Base62Chars
	}
	for i := 0; i < len(b); i++ {
		b[i] = letters[b[i]%byte(len(letters))]
	}

	return string(b)
}
