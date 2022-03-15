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

// http.Requestはアクセスした際の値が入っている // ResponseWriter wに対してresponseするものを返す
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// URLのpath情報が取れる
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	// nilを選択をするとデフォルトを返してくれる
	// ListenAndServe nilにするとデフォルトでPageNotFoundが出されるのでPageNotFoundが返される前にpathを指定をしないといけない
	log.Fatal(http.ListenAndServe(":8080", nil))
}
