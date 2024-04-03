// По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

// Входные данные
// Дано число n (0<n<100).

// Выходные данные
// Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

package main

import "fmt"

func main() {
	var n, first int
	var text string = "korov"
	fmt.Scan(&n)
	first = n % 10
	switch {
	case first == 1 && n != 11:
		text += "a"
	case first >= 2 && first <= 4 && (n < 12 || n > 14):
		text += "y"
	}
	fmt.Print(n, " ", text)
}
