// Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.

// Входные данные

// Программа получает на вход три натуральных числа: x, p, y.

// Выходные данные

// Программа должна вывести одно целое число.

// Sample Input:
// 100
// 10
// 200

// Sample Output:
// 8

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, p float64
	var years int
	fmt.Scan(&x, &p, &y)
	for ; x <= y; x += math.Round(float64(x)) / 100 * p {
		years++
	}
	fmt.Println(years)
}
