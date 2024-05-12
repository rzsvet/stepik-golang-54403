/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
Sample Input: zaabcbd
Sample Output: zcd
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	fmt.Scan(&text)
	for _, l := range text {
		if strings.Count(text, string(l)) == 1 {
			fmt.Print(string(l))
		}
	}
}
