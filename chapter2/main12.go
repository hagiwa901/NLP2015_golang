/*
各行の1列目だけを抜き出したものをcol1.txtに，2列目だけを抜き出したものをcol2.txtとしてファイルに保存せよ．
確認にはcutコマンドを用いよ．
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
	slice := strings.Split(str, "\n")

	var col1, col2 string

    // 出力
	for _, s := range slice {
		// 内容があるかどうか
		if(len(s) > 0){
			col := strings.Split(s, "\t")
			col1 += col[0] + "\n"
			col2 += col[1] + "\n"
		}
	}
	col1_byte := []byte(col1)
	col2_byte := []byte(col2)
	ioutil.WriteFile("chapter2/col1.txt", col1_byte, os.ModePerm)
	ioutil.WriteFile("chapter2/col2.txt", col2_byte, os.ModePerm)
}
