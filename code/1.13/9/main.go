// Количество минимумов
// Найдите количество минимальных элементов в последовательности.

// Входные данные

// Вводится натуральное число N, а затем N целых чисел последовательности.

// Выходные данные

// Выведите количество минимальных элементов последовательности.

// Sample Input:
// 3
// 21
// 11
// 4

// Sample Output:
// 1

package main

import "fmt"

func main() {

	var count, input, min, nums int

	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&input)
		if i == 0 {
			min = input
		}
		if input < min {
			min = input
			nums = 0
		}
		if input == min {
			nums++
		}
	}
	fmt.Println(nums)
}
