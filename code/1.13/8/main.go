// Количество нулей
// По данным числам, определите количество чисел, которые равны нулю.

// Входные данные
// Вводится натуральное число N, а затем N чисел.

// Выходные данные
// Выведите количество чисел, которые равны нулю.

// Sample Input:
// 5
// 1
// 8
// 100
// 0
// 12

// Sample Output:
// 1

package main

import "fmt"

func main() {
	var count, input, zero int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&input)
		if input == 0 {
			zero++
		}
	}
	fmt.Print(zero)
}
