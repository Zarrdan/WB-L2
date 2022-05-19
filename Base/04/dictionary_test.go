package main

import (
	"fmt"
	"testing"
)

func TestFindAnagrams(t *testing.T) {

	dictionary := []string{"Пятак", "пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	res := setAnagram(dictionary)
	fmt.Println(res)
	expected := map[string][]string{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}}
	fmt.Println(expected)
}
