// Требуется определить, является ли данный год високосным, напомним:
// Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
// - кратен 400;
// - кратен 4, но не кратен 100.

// Входные данные

// Вводится единственное число - номер года (целое, положительное, не превышает 10000).

// Выходные данные

// Требуется вывести слово YES, если год является високосным и NO - в противном случае.

// Sample Input:
// 2000

// Sample Output:
// YES

package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	if input%400 == 0 || input%4 == 0 && input%100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
