// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

// Входные данные
// Вводится натуральное число.

// Выходные данные
// Выведите ответ на задачу.

// Sample Input:
// 50

// Sample Output:
// 1 2 4 8 16 32

package main

import (
	"fmt"
	"math"
)

func main() {
	var N, result int
	fmt.Scan(&N)
	for i := 0; true; i++ {
		result = int(math.Pow(2, float64(i)))
		if result > N {
			break
		}
		fmt.Print(result, " ")
	}
}
