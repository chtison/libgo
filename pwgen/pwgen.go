package pwgen

import (
	"crypto/rand"
	"math/big"
)

// Constant predefined charsets.
const (
	CharsetLowercase   = `abcdefghijklmnopqrstuvwxyz`
	CharsetUppercase   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	CharsetNumeric     = `0123456789`
	CharsetPunctuation = `!"#$%&'()*+,-./:;<=>?@[\]^_{|}` + "`"
)

// Generate a string with n runes from charset
// Panics if len([]rune(charset)) == 0 and n != 0
func Generate(n uint, charset string) string {
	buffer := make([]rune, 0, n)
	runeset := []rune(charset)
	max := big.NewInt(int64(len(runeset)))
	for i := uint(0); i < n; i++ {
		i, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		c := runeset[i.Int64()]
		buffer = append(buffer, c)
	}
	return string(buffer)
}
