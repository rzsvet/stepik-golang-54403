// Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

// Формат входных данных

// На вход подается номер билета - одно шестизначное  число.

// Формат выходных данных
// Выведите "YES", если билет счастливый, в противном случае - "NO".

// Sample Input:
// 613244

// Sample Output:
// YES

package main

import "fmt"

func main() {
	var input, count int
	var nums [2]int

	for {
		_, err := fmt.Scanf("%1d", &input)

		if err != nil {
			break
		}
		if count < 3 {
			nums[0] += input
		} else {
			nums[1] += input
		}
		count += 1
	}
	if nums[0] == nums[1] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
