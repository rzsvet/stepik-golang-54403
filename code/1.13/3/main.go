// Дано трехзначное число. Переверните его, а затем выведите.

// Формат входных данных
// На вход дается трехзначное число, не оканчивающееся на ноль.

// Формат выходных данных
// Выведите перевернутое число.

// Sample Input:
// 843

// Sample Output:
// 348

package main

import (
	"fmt"
)

func main() {
	var input int
	var output string
	for {
		_, err := fmt.Scanf("%1d", &input)
		if err != nil {
			fmt.Print(output)
			break
		}
		output = fmt.Sprintf("%v", input) + output
	}
}
