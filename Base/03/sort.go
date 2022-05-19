package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	k       int
	n, r, u bool
}

func newFlags(k *int, n, r, u *bool) flags {
	return flags{
		k: *k,
		n: *n,
		r: *r,
		u: *u,
	}
}

func main() {
	k := flag.Int("k", -1, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	// структура с флагами
	flags := newFlags(k, n, r, u)
	// считываем файл.
	filePath := os.Args[len(os.Args)-1]
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	flags.sort(content)

}

func (f flags) sort(c []byte) {
	str := strings.Split(string(c), "\n")

	if f.u {
		temp := map[string]struct{}{}
		for _, r := range str {
			temp[r] = struct{}{}
		}
		str = make([]string, 0, len(temp))
		for k := range temp {
			str = append(str, k)
		}
	}
	dataArr := make([][]string, len(str))
	for i := range str {
		dataArr[i] = strings.Split(str[i], " ")
	}
	if f.k > -1 {
		switch true {
		case f.n && f.r:
			sortByNumDesc(&f, dataArr)
		case f.n:
			sortByNumAsc(&f, dataArr)
		case f.r:
			sortByLengthDesc(&f, dataArr)
		default:
			sortByLengthAsc(&f, dataArr)
		}
		fmt.Println(dataArr)
	}
	if f.k == -1 {
		switch true {
		case f.n && f.r:
			sortFullByNumDesc(&f, str)
		case f.n:
			sortFullByNumAsc(&f, str)
		case f.r:
			sortFullByLengthDesc(&f, str)
		default:
			sortFullByLengthAsc(&f, str)
		}
		fmt.Println(str)
	}

}

// Сортировка по возрастанию, с учетом колонки по длине строки
func sortByLengthAsc(f *flags, dataArr [][]string) {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return len(dataArr[j][f.k]) < len(dataArr[i][f.k])
	})
}

// Сортировка по уменьшению, с учетом колонки по длине строки
func sortByLengthDesc(f *flags, dataArr [][]string) {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return len(dataArr[j][f.k]) > len(dataArr[i][f.k])
	})
}

// Сортировка по уменьшению, с учетом колонки по длине числовому значению
func sortByNumAsc(f *flags, dataArr [][]string) {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return dataArr[j][f.k] > dataArr[i][f.k]
	})
}

// Сортировка по уменьшению, с учетом колонки по числовому значению
func sortByNumDesc(f *flags, dataArr [][]string) {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return dataArr[j][f.k] < dataArr[i][f.k]
	})
}

// Сортировка по возрастанию, по длине строки
func sortFullByLengthAsc(f *flags, dataArr []string) {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return len(dataArr[j]) < len(dataArr[i])
	})
}

// Сортировка по уменьшению,  по длине строки
func sortFullByLengthDesc(f *flags, dataArr []string) {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return len(dataArr[j]) > len(dataArr[i])
	})
}

// Сортировка по уменьшению, по длине числовому значению
func sortFullByNumAsc(f *flags, dataArr []string) {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return dataArr[j] > dataArr[i]
	})
}

// Сортировка по уменьшению, по числовому значению
func sortFullByNumDesc(f *flags, dataArr []string) {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return dataArr[j] < dataArr[i]
	})
}
