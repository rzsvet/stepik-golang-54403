/*
На вход подается строка.
Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input: топот
Sample Output: Палиндром
*/

package main

import (
	"fmt"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	var text string
	fmt.Scan(&text)

	if text == Reverse(text) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
