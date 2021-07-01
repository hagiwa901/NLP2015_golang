/*
Wikipedia記事のJSONファイルを読み込み，「イギリス」に関する記事本文を表示せよ．
問題21-29では，ここで抽出した記事本文に対して実行せよ．
*/

package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

type jawikiStruct struct {
	Text	string	`json: "text"`
    Title	string	`json: "title"`
}

func main() {
	/*
	https://qiita.com/nayuneko/items/2ec20ba69804e8bf7ca3
	*/

	// JSONファイル読み込み
	f, err := os.Open("chapter3/jawiki-country.json")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()
	
	b, err := ioutil.ReadFile("chapter3/jawiki-country.json")

	if err != nil {
		fmt.Println("json input error")
	}

	// JSONデコード
    var jawiki []jawikiStruct
    if err := json.Unmarshal(b, &jawiki); err != nil {
		if err, ok := err.(*json.SyntaxError); ok {
			fmt.Println(string(b[err.Offset-15:err.Offset+15]))
		}
        fmt.Println(err)
    }

	// デコードしたデータを表示
    for _, j := range jawiki {
		if(j.Title == "イギリス"){
			fmt.Printf("%s", j.Title)
		}
    }
}
