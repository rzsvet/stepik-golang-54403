// Номер числа Фибоначчи
// Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

// Входные данные
// Вводится натуральное число.

// Выходные данные
// Выведите ответ на задачу.

// Sample Input:
// 8

// Sample Output:
// 6

package main

import (
	"fmt"
)

func main() {
	var A int
	fmt.Scan(&A)
	prev, curr := 0, 1
	count := 1
	for curr < A {
		prev, curr = curr, prev+curr
		count++
	}
	if curr == A {
		fmt.Print(count)
	} else {
		fmt.Print(-1)
	}

}
