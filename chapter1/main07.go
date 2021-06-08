package main

import (
	"fmt"
	"strconv"
)

func template(x int, y string, z float64) string {
	var itoa string = strconv.Itoa(x)
	var float string = strconv.FormatFloat(z, 'f',  1, 64)
	return itoa + "時の" + y + "は" + float
}

func main() {
	fmt.Println(template(12, "気温", 22.4))
}
