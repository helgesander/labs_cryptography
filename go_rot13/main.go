package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	result_test := rot13(text)
	fmt.Println(result_test)
}

func rot13(txt string) string {
	result := make([]byte, len(txt), len(txt))
	for i := 0; i < len(txt); i++ {
		tmp := txt[i]
		if tmp >= 'A' && tmp <= 'z' {
			tmp += 13
		}
		if tmp > 'z' {
			tmp -= 26
		}
		result[i] = tmp
	}
	return string(result)
}
