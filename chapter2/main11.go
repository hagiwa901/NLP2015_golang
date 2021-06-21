/*
タブ1文字につきスペース1文字に置換せよ．
確認にはsedコマンド，trコマンド，もしくはexpandコマンドを用いよ．
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

	str := string(b)
	str = strings.Replace(str, "\t", " ", -1)
	slice := strings.Split(str, "\n")

    // 出力
	for _, s := range slice {
		// 内容があるかどうか
		if(len(s) > 4){
			fmt.Println(s)
		}
	}
}
