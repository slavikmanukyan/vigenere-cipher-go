package main

import (
	"fmt"
)

func cipher(s string, key string) string {
	result := []rune(s)
	keyLen := len(key)
	spaceCount := 0
	for index, char := range result {
		if char >= 'a' && char <= 'z' {
			newChar := (s[index]+key[(index-spaceCount)%keyLen])%97 + 97
			if newChar >= 97+26 {
				newChar = newChar - 26
			}
			result[index] = rune(newChar)
		}
		if char == ' ' {
			spaceCount++
		}
	}
	return string(result)
}

func main() {
	fmt.Println(cipher("dasjkdhas djkasd asjkdahsd ajksd asdjkahsd asjkdhasd asdiasdhajksdhjasd", "xaxkj"))
}
