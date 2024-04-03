// Из натурального числа удалить заданную цифру.

// Входные данные
// Вводятся натуральное число и цифра, которую нужно удалить.

// Выходные данные
// Вывести число без заданных цифр.

// Sample Input:
// 38012732
// 3

// Sample Output:
// 801272

package main

import (
	"fmt"
)

func main() {
	var input, last_number, number int
	var output string
	fmt.Scan(&input, &number)
	for input != 0 {
		last_number = input % 10
		if last_number != number {
			output = fmt.Sprintf("%v", last_number) + output
		}
		input = input / 10
	}
	fmt.Print(output)
}
