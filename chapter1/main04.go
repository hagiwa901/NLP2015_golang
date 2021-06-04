package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	var s1 string
	// ,の削除
	str = strings.Replace(str, ",", "", -1)
	// .の削除
	str = strings.Replace(str, ".", "", -1)
	slice := strings.Split(str, " ")
	for _, s := range slice {
		s1 += s[0:1]
	}
	fmt.Println(s1)
}
