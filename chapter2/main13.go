/*
12で作ったcol1.txtとcol2.txtを結合し，元のファイルの1列目と2列目をタブ区切りで並べたテキストファイルを作成せよ．
確認にはpasteコマンドを用いよ．
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
    f1, err := os.Open("chapter2/col1.txt")
    if err != nil{
        fmt.Println("error")
    }
    defer f1.Close()

    // 一気に全部読み取り
    b1, err := ioutil.ReadAll(f1)

	str1 := string(b1)
	slice1 := strings.Split(str1, "\n")

	// ファイルをOpenする
    f2, err := os.Open("chapter2/col2.txt")
    if err != nil{
        fmt.Println("error")
    }
    defer f2.Close()

    // 一気に全部読み取り
    b2, err := ioutil.ReadAll(f2)

	str2 := string(b2)
	slice2 := strings.Split(str2, "\n")

    var col string
	for i := 0; i < len(slice1); i++ {
		col += slice1[i] + "\t" + slice2[i] + "\n"
	}
	fmt.Println(col)
}
