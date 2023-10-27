package strings

import (
	"unicode/utf8"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/stack"
)

func Reverse(input string) string {
	s := stack.Stack[rune]{}
	for _, r := range input {
		s.Push(r)
	}
	var output []byte
	for !s.IsEmpty() {
		nextRune := s.Pop()
		nextChar := make([]byte, utf8.RuneLen(*nextRune))
		utf8.EncodeRune(nextChar, *nextRune)
		output = append(output, nextChar...)
	}

	return string(output)
}
