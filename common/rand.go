package common

import (
	"math/rand"
	"strings"
	"time"
)

var (
	Char    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr = strings.Split(Char, "")
	charLen = len(charArr)
	seed    = rand.New(rand.NewSource(time.Now().Unix()))
)

// RandString randomly generate a string based on length
func RandString(length int) string {
	if length < 1 {
		return ""
	}

	var res = ""
	for i := 1; i <= length; i++ {
		res = res + charArr[seed.Intn(charLen)]
	}
	return res
}
