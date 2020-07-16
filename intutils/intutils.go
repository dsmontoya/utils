package intutils

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/dsmontoya/utils/internal"
)

// GroupDigits groups each n digits of a number from right to left. Use sep as the seperator for each group.
func GroupDigits(number int, sep string, n int) string {
	str := strconv.Itoa(number)
	return internal.StrGroupDigits(str, sep, n)
}

// RandInt returns a random integer between min and max
func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func RandUint(min uint, max uint) uint {
	return uint(RandInt(int(min), int(max)))
}

func Test() {

}
