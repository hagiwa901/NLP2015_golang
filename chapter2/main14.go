/*
自然数Nをコマンドライン引数などの手段で受け取り，入力のうち先頭のN行だけを表示せよ．
確認にはheadコマンドを用いよ
*/

package main

import(
    "fmt"
    "os"
    "io/ioutil"
	"strings"
	"flag"
)

func main() {
	// コマンドライン引数
    args := flag.Int("n", 24, "help message for n option")
	flag.Parse()

	// ファイルをOpenする
    f, err := os.Open("chapter2/hightemp.txt")
    if err != nil{
        fmt.Println("error")
    }
    defer f.Close()

    // 一気に全部読み取り
    b, err := ioutil.ReadAll(f)

	str := string(b)
	slice := strings.Split(str, "\n")

    // 出力
	for i := 0; i < *args; i++ {
		fmt.Println(slice[i])
	}
}
