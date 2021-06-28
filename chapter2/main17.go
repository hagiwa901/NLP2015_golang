/*
1列目の文字列の種類（異なる文字列の集合）を求めよ．
確認にはsort, uniqコマンドを用いよ．
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// ファイルをOpenする
	f, err := os.Open("chapter2/hightemp.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	// 一気に全部読み取り
	b, err := ioutil.ReadAll(f)

	str := string(b)
	slice := strings.Split(str, "\n")

	// 重複の削除
	m := map[string]bool{}

	// 出力
	for _, s := range(slice) {
		if(len(s) > 0){
			str := strings.Split(s, "\t")
			if !m[str[0]] {
				m[str[0]] = true
				fmt.Println(str[0])
			}
		}
	}
}
