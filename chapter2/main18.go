/*
各行を3コラム目の数値の逆順で整列せよ（注意: 各行の内容は変更せずに並び替えよ）．
確認にはsortコマンドを用いよ（この問題はコマンドで実行した時の結果と合わなくてもよい）．
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sort"
	"strconv"
)

type pair struct {
    key float64
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
	pairs := []pair{}

	// 出力
	for _, s := range(slice) {
		if(len(s) > 0){
			str := strings.Split(s, "\t")
			num, _ := strconv.ParseFloat(str[2], 64)
			pairs = append(pairs, pair{num, s})
		}
	}

	// 降順ソート
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].key > pairs[j].key
    })

	for i := 0; i < len(pairs); i++ {
		fmt.Println(pairs[i].str)
	}
}
