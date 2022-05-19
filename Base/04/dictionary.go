package main

/*
=== Поиск анаграмм по словарю ===
Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	testStr := []string{"Пятак", "пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	temp := setAnagram(testStr)
	for k, v := range temp {
		fmt.Println(k, ":", v)
	}
}

func setAnagram(str []string) map[string][]string {

	s := checkDubl(str)

	tempMap := make(map[string][]string)
	for _, val := range s {
		tempMap[strSort(val)] = append(tempMap[strSort(val)], val)
	}

	// мапа множест анаграм, у которой ключ первое встретившейся в словаре слово из множества.
	newMap := make(map[string][]string)
	for _, v := range tempMap {
		temp := v
		newMap[v[0]] = temp
	}
	return newMap
}

func strSort(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func checkDubl(s []string) []string {
	temp := map[string]struct{}{}
	for _, v := range s {
		temp[strings.ToLower(v)] = struct{}{}
	}
	str := make([]string, 0)
	for k := range temp {
		str = append(str, k)
	}
	return str
}
