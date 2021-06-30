/*
各行の1列目の文字列の出現頻度を求め，その高い順に並べて表示せよ．
確認にはcut, uniq, sortコマンドを用いよ
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sort"
)

type pairs struct {
	num int
	str string
}

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

	// Key
	m := map[string]int{}

	// 出力
	for _, s := range(slice) {
		if(len(s) > 0){
			str := strings.Split(s, "\t")
			m[str[0]] += 1
		}
	}

	p := []pairs{}

	for k, v := range m {
		p = append(p, pairs{v, k})
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].num > p[j].num
	})

	for i := 0; i < len(m); i++ {
		fmt.Println(p[i].str, p[i].num)
	}
}

