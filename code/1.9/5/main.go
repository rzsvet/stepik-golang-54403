// На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное", если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.

// Sample Input:
// 5

// Sample Output:
// Число положительное

package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)
	switch {
	case input == 0:
		fmt.Println("Ноль")
	case input > 0:
		fmt.Println("Число положительное")
	default:
		fmt.Println("Число отрицательное")
	}
}
