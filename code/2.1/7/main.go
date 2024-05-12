// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

// Входные данные

// Вводится четыре числа.

// Выходные данные

// Необходимо вернуть из функции наименьшее из 4-х данных чисел

// Sample Input:
// 4 5 6 7

// Sample Output:
// 4

func minimumFromFour() int {
	var min, input int
	for i := 0; i <= 3; i++ {
		fmt.Scan(&input)
		if i == 0 {
			min = input
		}
		if min > input {
			min = input
		}
	}
	return min
}