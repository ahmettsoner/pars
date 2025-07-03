package random

import (
	"math/rand"
	"strings"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt returns a random integer between min and max (inclusive).
func Int(min, max int) int {
	if max <= min {
		return min
	}
	return seededRand.Intn(max-min+1) + min
}

// RandomString returns a random string of given length using letters only.
func String(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return stringWithCharset(length, letters)
}

// RandomStringWithNumbers returns a random string including numbers.
func StringWithNumbers(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	return stringWithCharset(length, letters)
}

// RandomEmail returns a fake but valid-looking email address.
func Email() string {
	return String(8) + "@" + String(5) + ".com"
}

// RandomBool returns true or false randomly.
func Bool() bool {
	return seededRand.Intn(2) == 1
}

// RandomFloat returns a float64 in the range [min, max).
func Float(min, max float64) float64 {
	return min + seededRand.Float64()*(max-min)
}

// RandomChoice picks a random item from a slice of strings.
func Choice(items []string) string {
	if len(items) == 0 {
		return ""
	}
	return items[seededRand.Intn(len(items))]
}

// RandomWords returns a slice of random words (strings).
func Words(count, wordLength int) []string {
	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = String(wordLength)
	}
	return words
}

// internal helper
func stringWithCharset(length int, charset []rune) string {
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(charset[seededRand.Intn(len(charset))])
	}
	return b.String()
}
