// Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
// ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.

// Sample Input: 2 4
// Sample Output: 4 2

func test(x1 *int, x2 *int) {
	*x2, *x1 = *x1, *x2
	fmt.Println(*x1, *x2)
}