package main

import (
	"fmt"
	"strconv"
	"strings"
)

func update(s string) (str string, err error) {

	s = strings.ToLower(s)

	if s == "" || len(s) == 1 {
		return "", fmt.Errorf("Некорректная строка")
	}
	if 96 >= s[0] || s[0] >= 123 {
		return "", fmt.Errorf("Некорректная строка")
	}
	for i := 0; i <= len(s)-1; i++ {
		// символ в алфавите то плюсую к стр
		if s[i] >= 97 && s[i] <= 122 {
			str += string(s[i])
			// если есть escape последовательность, записываю в стр следующий символ, перемешаю индекс в массиве на 1 вправо
		} else if s[i] == '\\' {
			str += string(s[i+1])
			i++
			// расчет числа повторений за -1, 1 символ уже есть в стр
		} else {
			r := make([]byte, 0)
			for n := i; n <= len(s)-1; n++ {
				if s[n] >= 48 && s[n] <= 57 {
					r = append(r, s[n])
				}
				break
			}
			y := ""
			a, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			//y := strings.Repeat(string(s[i-1]), a-1)
			for n := i; 1 < a; a-- {
				y += string(s[n-1])
			}
			str += y
		}

	}
	return str, err
}

func main() {
	//str := "abc4d"
	str2, err := update("a4bc2d5e")
	fmt.Println(str2, "\n", err)

}
