/*
Wikipedia記事のJSONファイルを読み込み，「イギリス」に関する記事本文を表示せよ．
問題21-29では，ここで抽出した記事本文に対して実行せよ．
*/

// https://cipepser.hatenablog.com/entry/2017/03/17/231556

package main

import (
	"encoding/json"
	"bufio"
	"fmt"
 	"os"
	"io"
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

	r := bufio.NewReader(f)

	// JSONデコード
	var jawikis []jawikiStruct
	
	for {
		// fmt.Println(scanner)
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		var jawiki jawikiStruct
		json.Unmarshal([]byte(b), &jawiki)
		jawikis = append(jawikis, jawiki)
	}

	for _, jawiki := range jawikis {
		if(jawiki.Title == "イギリス") {
			fmt.Println(jawiki.Text)
		}
	}
	
	f.Close()
}
