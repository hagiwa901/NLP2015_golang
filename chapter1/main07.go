/*
引数x, y, zを受け取り「x時のyはz」という文字列を返す関数を実装せよ．
さらに，x=12, y="気温", z=22.4として，実行結果を確認せよ．
*/

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
