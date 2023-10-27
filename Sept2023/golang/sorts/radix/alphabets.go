package radix

import (
	"unicode/utf8"

	"golang.org/x/exp/slices"
)

type alphabet rune

const (
	A     alphabet = 'A'
	B     alphabet = 'B'
	C     alphabet = 'C'
	D     alphabet = 'D'
	E     alphabet = 'E'
	F     alphabet = 'F'
	G     alphabet = 'G'
	H     alphabet = 'H'
	I     alphabet = 'I'
	J     alphabet = 'J'
	K     alphabet = 'K'
	L     alphabet = 'L'
	M     alphabet = 'M'
	N     alphabet = 'N'
	O     alphabet = 'O'
	P     alphabet = 'P'
	Q     alphabet = 'Q'
	R     alphabet = 'R'
	S     alphabet = 'S'
	T     alphabet = 'T'
	U     alphabet = 'U'
	V     alphabet = 'V'
	W     alphabet = 'W'
	X     alphabet = 'X'
	Y     alphabet = 'Y'
	Z     alphabet = 'Z'
	a     alphabet = 'a'
	b     alphabet = 'b'
	c     alphabet = 'c'
	d     alphabet = 'd'
	e     alphabet = 'e'
	f     alphabet = 'f'
	g     alphabet = 'g'
	h     alphabet = 'h'
	i     alphabet = 'i'
	j     alphabet = 'j'
	k     alphabet = 'k'
	l     alphabet = 'l'
	m     alphabet = 'm'
	n     alphabet = 'n'
	o     alphabet = 'o'
	p     alphabet = 'p'
	q     alphabet = 'q'
	r     alphabet = 'r'
	s     alphabet = 's'
	t     alphabet = 't'
	u     alphabet = 'u'
	v     alphabet = 'v'
	w     alphabet = 'w'
	x     alphabet = 'x'
	y     alphabet = 'y'
	z     alphabet = 'z'
	zero  alphabet = '0'
	one   alphabet = '1'
	two   alphabet = '2'
	three alphabet = '3'
	four  alphabet = '4'
	five  alphabet = '5'
	six   alphabet = '6'
	seven alphabet = '7'
	eight alphabet = '8'
	nine  alphabet = '9'
)

type Alphabet interface {
	Alphabet() alphabet
	Rune() rune
}

func (a alphabet) Alphabet() alphabet {
	return a
}

func (a alphabet) Rune() rune {
	return rune(a)
}

func (a alphabet) String() string {
	len := utf8.RuneLen(a.Rune())
	b := make([]byte, len)
	utf8.EncodeRune(b, a.Rune())
	return string(b)
}

var lexicographicallySortedAlphabetsMap map[alphabet]int

func LexicographicallySortedAlphabetsMap() map[alphabet]int {
	if lexicographicallySortedAlphabetsMap != nil {
		return lexicographicallySortedAlphabetsMap
	}
	alphabets := []alphabet{A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z, zero, one, two, three, four, five, six, seven, eight, nine}
	slices.SortFunc(alphabets, func(a alphabet, b alphabet) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	lexicographicallySortedAlphabetsMap := make(map[alphabet]int)
	for i, r := range alphabets {
		lexicographicallySortedAlphabetsMap[r] = i + 1
	}

	return lexicographicallySortedAlphabetsMap
}
