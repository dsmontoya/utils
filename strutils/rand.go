package strutils

import (
	"github.com/dsmontoya/utils/intutils"
)

var allChars = append(alphaNum, symbols...)
var alphaNum = append(letters, numbers...)
var letters = append(lowercaseLetters, uppercaseLetters...)
var numbers = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var symbols = []byte{'!', '#', '$', '%', '&', '(', ')', '*', '+', '.', '/', '<', '=', '>', ';', '?', '@', '[', ']', '_', '{', '}', '|', '~', ' ', 'ñ'}
var lowercaseLetters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var uppercaseLetters = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

//Rand generates a random string with l length
func Rand(l int) string {
	s := make([]byte, l)
	for i := 0; i < l; i++ {
		n := intutils.RandUint(0, uint(len(allChars)-1))
		s[i] = allChars[n]
	}
	return string(s)
}
