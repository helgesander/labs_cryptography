package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("dsfsfs")
	text, _ := in.ReadString('\n')
	for i := 1; i < 26; i++ {
		fmt.Printf("%d: %s\n", i, brute_caesar(strings.Trim(text, "\r\n"), i))
	}
}

func brute_caesar(txt string, key int) string {
	result := make([]byte, len(txt))
	const alphabet_upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const alphabet_lower string = "abcdefghijklmnopqrstuvwxyz"
	for idx, chr := range txt {
		tmp := txt[idx]
		if tmp >= 'A' && tmp <= 'z' {
			if unicode.IsUpper(chr) {
				result[idx] = alphabet_upper[(tmp-65+byte(key))%26]
			} else {
				result[idx] = alphabet_lower[(tmp-97+byte(key))%26]
			}
		} else {
			result[idx] = tmp
		}
	}
	return string(result)
}
