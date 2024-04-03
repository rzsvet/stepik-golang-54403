// Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.

// Формат входных данных

// Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит, а служит как признак ее окончания).

// Формат выходных данных

// Выведите ответ на задачу.

// Sample Input:
// 1
// 3
// 3
// 1
// 0

// Sample Output:
// 2

package main

import "fmt"

func main() {

	var input int
	var nums [100]int
	max := 0

	for fmt.Scan(&input); input != 0; fmt.Scan(&input) {
		if input > max {
			max = input
		}
		nums[input]++
	}
	fmt.Println(nums[max])
}
