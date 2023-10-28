package radix

import (
	"fmt"
	"unicode/utf8"
)

type runeArray struct {
	runes []rune
}

func newRuneArray(s string) runeArray {
	runes := []rune{}
	for _, r := range s {
		runes = append(runes, r)
	}
	return runeArray{
		runes: runes,
	}
}

func (r *runeArray) RuneAt(i int) rune {
	return r.runes[i]
}

func (r runeArray) String() string {
	bArr := []byte{}
	for _, ru := range r.runes {
		size := utf8.RuneLen(ru)
		b := make([]byte, size)
		utf8.EncodeRune(b, ru)
		bArr = append(bArr, b...)
	}
	return string(bArr)
}

func (r *runeArray) Len() int {
	return len(r.runes)
}

type ArrayOfRuneArray []runeArray

func (a ArrayOfRuneArray) String() string {
	str := ""
	for _, v := range a {
		str = fmt.Sprintf("%s %s", str, runeArray(v))
	}
	return str
}
