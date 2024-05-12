/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input: ihgewlqlkot
Sample Output: hello
*/

package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	for i, _ := range text {
		if i%2 != 0 {
			fmt.Print(string(text[i]))
		}
	}
}
