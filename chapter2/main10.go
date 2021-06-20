/*
行数をカウントせよ．確認にはwcコマンドを用いよ．

% wc chapter2/hightemp.txt
      24      98     813 chapter2/hightemp.txt
*/
package main

import(
    "fmt"
    "os"
    "io/ioutil"
	"strings"
)

func main() {
	// ファイルをOpenする
    f, err := os.Open("chapter2/hightemp.txt")
    if err != nil{
        fmt.Println("error")
    }
    defer f.Close()

    // 一気に全部読み取り
    b, err := ioutil.ReadAll(f)

	// 行数
	var column = 0

	str := string(b)
	slice := strings.Split(str, "\n")

    // 出力
	for _, s := range slice {
		// 内容があるかどうか
		if(len(s) > 0){
			column++
		}
	}
    fmt.Println(column)
}
