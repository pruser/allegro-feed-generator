package utils

import (
	"strings"
	"unicode/utf8"
)

func Cover(str string, showOnly uint) string {
	len := uint(utf8.RuneCountInString(str))
	if showOnly >= len {
		return str
	}
	head := []rune(str)[:showOnly]
	tail := []rune(strings.Repeat("*", int(len-showOnly)))

	return string(append([]rune(head), tail...))
}
