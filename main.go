package main

import (
	"fmt"
)

func cipher(s string, key string, encrypt bool) string {
	result := []rune(s)
	keyLen := len(key)
	spaceCount := 0
	for index, char := range result {
		newChar := byte(char)
		m1 := byte('A')
		m2 := byte('A')
		if char >= 'a' {
			m1 = byte('a')
		}
		keyChar := key[(index-spaceCount)%keyLen]
		if keyChar >= 'a' {
			m2 = byte('a')
		}
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if encrypt {
				newChar = (s[index]-m1+(keyChar-m2))%26 + m1
			} else {
				newChar = (s[index]-m1-(keyChar-m2)+26)%26 + m1
			}

			result[index] = rune(newChar)
		} else {
			spaceCount++
		}
	}
	return string(result)
}

func encrypt(in string, key string) string {
	return cipher(in, key, true)
}

func decrypt(in string, key string) string {
	return cipher(in, key, false)
}

func main() {
	t := encrypt("dasjkdhas djkasd asjkdahsd ajksd asdjkahsd asjkdhasd asdiasdhajksdhjasd", "xaxkj")
	s := decrypt(t, "xaxkj")
	fmt.Println(t)
	fmt.Println(s)
}
