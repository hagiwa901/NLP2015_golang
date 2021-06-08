package main

import (
	"github.com/deckarep/golang-set"
	"fmt"
)

func ngram(str string, n int) []string {
	// 文字
	size := len(str)
	var gram []string
	for i := 0; i < size; i = i+1 {
		if i + n <= size {
			gram = append(gram, str[i:i+n])
		}
	}
	return gram
}

func main() {
	str1Set := mapset.NewSet()
	str2Set := mapset.NewSet()
	const n int = 2
	str1 := ngram("paraparaparadise", n)
	str2 := ngram("paragraph", n)

	fmt.Println(str1)
	fmt.Println(str2)

	for _, s := range str1 {
		str1Set.Add(s)
	}
	
	for _, s := range str2 {
		str2Set.Add(s)
	}
	// 和集合
	fmt.Println(str1Set.Union(str2Set))
	// 積集合
	fmt.Println(str1Set.Intersect(str2Set))
	// 差集合
	fmt.Println(str1Set.Difference(str2Set).Union(str2Set.Difference(str1Set)))
	// [se]が含まれるかどうか
	fmt.Println(str1Set.Contains("se"))
	fmt.Println(str2Set.Contains("se"))
}
