/*
自然数Nをコマンドライン引数などの手段で受け取り，入力のファイルを行単位でN分割せよ．
同様の処理をsplitコマンドで実現せよ．
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// コマンドライン引数
	args := flag.Int("n", 1, "help message for n option")
	flag.Parse()

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

	// 分割点
	SplitNum := len(slice) / *args

	num := make([]bool, len(slice))

	for i := 0; i < len(slice); i++ {
		if((i+1) % SplitNum == 0){
			num[i] = true
		} else {
			num[i] = false
		}
	}

	// 出力
	for i, s := range(slice) {
		if(len(s) > 0){
			if(num[i]){
				fmt.Println(s)
				fmt.Println()
			}else{
				fmt.Println(s)
			}
		}
	}
}

