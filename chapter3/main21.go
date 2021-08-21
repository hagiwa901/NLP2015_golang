/*
記事中でカテゴリ名を宣言している行を抽出せよ．
*/

// https://cipepser.hatenablog.com/entry/2017/03/17/231556

package main

import (
	"encoding/json"
	"bufio"
	"fmt"
 	"os"
	"io"
	"regexp"
)

type Article struct {
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

	// 1行ずつ処理
	r := bufio.NewReader(f)

	// JSONデコード
	var jawiki []Article
	
	for {
		// fmt.Println(scanner)
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		var article Article
		json.Unmarshal([]byte(b), &article)
		jawiki = append(jawiki, article)
	}

	rep := regexp.MustCompile(`\[\[Category:.+\]\]`)
	for _, j := range jawiki {
		res := rep.FindAllStringSubmatch(j.Text, -1)
		for _, v := range res {
			fmt.Println(v)
		}
	}
	
	f.Close()
}
