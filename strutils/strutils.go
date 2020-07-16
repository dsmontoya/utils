package strutils

import (
	"fmt"
	"regexp"
	"unicode"
	"unicode/utf8"

	"github.com/dsmontoya/utils/internal"
)

// Copy makes a copy of a string
func Copy(str string) string {
	b := str
	b2 := make([]byte, len(b))
	copy(b2, b)
	return string(b2)
}

//Urif formats the uri parameters with the format :resource_id
func Urif(uri string, values ...interface{}) string {
	r := regexp.MustCompile(`:([A-z]|[0-9]|_|)+`)
	f := r.ReplaceAll([]byte(uri), []byte("%v"))
	return fmt.Sprintf(string(f), values...)
}

// GroupDigits groups each n digits of a number from right to left. Use sep as the seperator for each group.
func GroupDigits(str, sep string, n int) string {
	return internal.StrGroupDigits(str, sep, n)
}

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

func (b *buffer) indent() {
	if len(b.r) > 0 {
		b.r = append(b.r, '_')
	}
}

func ToSnakeCase(s string) string {
	b := buffer{
		r: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.indent()
					w = true
				}
				b.write(m)
			}
			m = unicode.ToLower(ch)
		} else {
			if m != 0 {
				b.indent()
				b.write(m)
				m = 0
				w = false
			}
			b.write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.indent()
		}
		b.write(m)
	}
	return string(b.r)
}

func PadLeft(source string, char string, length int) string {

	if len(source) < length {
		complete := ""
		for i := 0; i < length-len(source); i++ {
			complete += char
		}

		source = complete + source
	}

	return source
}

func DecodeUTF8(source string) string {
	res := ""
	for len(source) > 0 {
		r, size := utf8.DecodeRuneInString(source)
		res += fmt.Sprintf("%c", r)
		source = source[size:]
	}
	return res
}
