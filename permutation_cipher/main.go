package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	ciphertext := permutation_encode(text)
	fmt.Printf("Encoded: %s", ciphertext)
	fmt.Printf("Decoded: %s", permutation_decode(ciphertext))
}

func permutation_encode(text string) string {
	const cipher_alphabet string = "HOMGIZYTPRQADBCJFEKNLSUVWXhomgizytprqadbcjfeknlsuvwx"
	result := make([]byte, len(text))
	for idx, chr := range text {
		tmp := text[idx]
		if tmp >= 'A' && tmp <= 'z' {
			if unicode.IsUpper(chr) {
				result[idx] = cipher_alphabet[tmp-65]
			} else {
				result[idx] = cipher_alphabet[tmp-71]
			}
		} else {
			result[idx] = tmp
		}
	}
	return string(result)
}

func permutation_decode(text string) string {
	const alphabet string = "LNOMRQDAEPSUCTBIKJVHWXYZGFlnomrqdaepsuctbikjvhwxyzgf"
	result := make([]byte, len(text))
	for idx, chr := range text {
		tmp := text[idx]
		if tmp >= 'A' && tmp <= 'z' {
			if unicode.IsUpper(chr) {
				result[idx] = alphabet[tmp-65]
			} else {
				result[idx] = alphabet[tmp-71]
			}
		} else {
			result[idx] = tmp
		}
	}
	return string(result)
}
