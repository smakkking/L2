package main

import (
	"fmt"
	"sort"
	"unicode"
)

type Key [33]rune

func ToLowerWord(word string) string {
	test := []rune(word)
	for i := 0; i < len(test); i++ {
		test[i] = unicode.ToLower(test[i])
	}
	return string(test)
}

func strKey(str string) Key {
	var key Key
	for _, elem := range []rune(str) {
		key[unicode.ToLower(elem)-'а']++
	}
	return key
}

func GroupAnagrams(strs []string) map[string][]string {
	groups := make(map[Key][]string)

	first_apperance := make(map[Key]string)

	for _, v := range strs {
		key := strKey(v)
		_, ok := first_apperance[key]
		if !ok {
			first_apperance[key] = ToLowerWord(v)
		}
		groups[key] = append(groups[key], ToLowerWord(v))
	}

	result := make(map[string][]string)

	for key := range first_apperance {
		if len(groups[key]) == 1 {
			continue
		}

		sort.Strings(groups[key])
		result[first_apperance[key]] = groups[key]
	}
	return result
}

func main() {
	//arr := [][]rune{[]rune(`пятка`), []rune(`тяпка`), []rune(`лом`), []rune(`мол`)}

	arr1 := []string{`Пятка`, `тяпка`, `лом`, `мол`}

	fmt.Println(GroupAnagrams(arr1))
}
