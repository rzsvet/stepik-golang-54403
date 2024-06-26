/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

Входные данные
Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

Выходные данные
Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input: 1112221112
Sample Output: 2
*/

package main

import (
	"fmt"
)

func main() {
	var input, max string
	fmt.Scan(&input)
	for _, num := range []rune(input) {
		if max < string(num) {
			max = string(num)
		}
	}
	fmt.Println(max)
}
