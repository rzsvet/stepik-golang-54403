/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры и буквы латинского алфавита.
На вход подается строка-пароль.
Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input: fdsghdfgjsdDD1
Sample Output: Ok
*/

package main

import (
	"fmt"
	"unicode"
)

func IsValidPassword(password string) bool {
	rune_password := []rune(password)
	if len(rune_password) < 5 {
		return false
	}

	for _, l := range rune_password {
		if !unicode.Is(unicode.Latin, l) && !unicode.Is(unicode.Digit, l) {
			return false
		}
	}

	return true
}

func main() {
	var password string
	fmt.Scan(&password)

	if IsValidPassword(password) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
