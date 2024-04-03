// Дано трехзначное число. Найдите сумму его цифр.

// Формат входных данных
// На вход дается трехзначное число.

// Формат выходных данных
// Выведите одно целое число - сумму цифр введенного числа.

// Sample Input:
// 745

// Sample Output:
// 16

package main

import (
	"fmt"
)

func main() {
	var input, sum int
	for {
		_, err := fmt.Scanf("%1d", &input)
		if err != nil {
			fmt.Print(sum)
			break
		}
		sum += input
	}
}
