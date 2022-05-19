package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fields      string
	delimiter   string
	isSeparated bool
)

func main() {
	flag.StringVar(&fields, "f", "", "Выбрать поля (колонки). Перечислить значения через запятую")
	flag.StringVar(&delimiter, "d", "\t", "Использовать другой разделитель")
	flag.BoolVar(&isSeparated, "s", false, "Выводить только строки c разделителем")

	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)

	data := [][]string{}
	for {
		fmt.Print("Введите строку или нажмите enter, чтобы завершить ввод: ")
		ok := scanner.Scan()
		if !ok && scanner.Err() == nil {
			break
		}
		str := scanner.Text()
		if len(str) == 0 {
			break
		}
		if isSeparated {
			if !strings.Contains(str, delimiter) {
				continue
			}
		}
		data = append(data, strings.Split(str, delimiter))
	}
	if fields != "" {
		numsFields := strings.Split(fields, ",")
		for _, i := range numsFields {
			num, err := strconv.Atoi(string(i))
			if err != nil {
				log.Println(err)
			}
			if num-1 >= len(data) {
				continue
			}
			fmt.Println(data[num-1])
		}
		return
	}
	fmt.Println(data)
}
