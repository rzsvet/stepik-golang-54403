// Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.

// Формат выходных данных
// Выведите одно целое число - первую цифру заданного числа.

// Sample Input:
// 1234

// Sample Output:
// 1

package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%1d", &input)
	fmt.Print(input)
}
