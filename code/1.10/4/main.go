// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

// Sample Input:
// 5
// 38 24 800 8 16

// Sample Output:
// 40

package main

import "fmt"

func main() {

	var count, input int
	fmt.Scan(&count)
	sum := 0
	for i := 1; i <= count; i++ {
		fmt.Scan(&input)
		if input%8 == 0 && input < 100 && input > 10 {
			sum += input
		}
	}
	fmt.Println(sum)
}
