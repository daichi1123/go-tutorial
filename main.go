package main

import (
	"fmt"
	"io/ioutil"
)

// Pageの管理をする構造体を作成する
type Page struct {
	Title string
	Body  []byte
}

// こちらを実行するとTitleに入れた文字列で.txtファイルを作成してBodyの中に入れた値がファイル内に記述される

// 構造体をpとして使用して返り値はerrorにしている
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 呼び込むためのページの作成 返り値がpage&error
func loadPage(title string) (*Page, error) {
	// どのファイルを読み込むかを変数に入れ
	filename := title + ".txt"
	// その特定したファイルを読み込む
	body, err := ioutil.ReadFile(filename)
	// errがあった場合のエラーハンドリング
	if err != nil {
		return nil, err
	}
	// errがなければ
	return &Page{Title: title, Body: body}, nil
}

func main() {
	// p1で構造体の情報を記述していくアドレスを指定してそのアドレスの値を定義している
	p1 := &Page{Title: "test", Body: []byte("this is a sample Page")}
	// そしてそのp1に対してsaveメソッドを使用する
	p1.save()

	// databaseに入っている中身をloadPageで取ってくる
	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}
