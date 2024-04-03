// По данному трехзначному числу определите, все ли его цифры различны.

// Формат входных данных
// На вход подается одно натуральное трехзначное число.

// Формат выходных данных
// Выведите "YES", если все цифры числа различны, в противном случае - "NO".

// Sample Input 1:
// 237

// Sample Output 1:
// YES

// Sample Input 2:
// 117

// Sample Output 2:
// NO

package main

import "fmt"

func main() {
	var input int
	var nums [10]bool

	for {
		_, err := fmt.Scanf("%1d", &input)

		if err != nil {
			fmt.Println("YES")
			break
		} else if nums[input] {
			fmt.Println("NO")
			break
		}
		nums[input] = true
	}
}
