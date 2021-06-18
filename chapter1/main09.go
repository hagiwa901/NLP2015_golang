package main

import (
	"strings"
	"fmt"
	"math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func randString(s string) string {
	if(len(s) > 4) {
		var randStr string
		randStr += s[0:1]
		for i := 1; i < len(s) - 1; i++ {
			randNum := rand.Intn(len(s))
			randStr += s[randNum:randNum+1]
		}
		randStr += s[len(s)-1:len(s)]
		return randStr
	} else {
		return s
	}
}

func cipher(str string) string {
	slice := strings.Split(str, " ")
	var rStr string
	for _, s := range slice {
		if(len(s) > 4) {
			rStr += randString(s)
		} else {
			rStr += s
		}
		rStr += " "
	}
	return rStr
}

func main() {
	str := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	fmt.Println(cipher(str))
}
