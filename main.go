package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// Pageの管理をする構造体を作成する
type Page struct {
	Title string
	Body  []byte
}

// Configで使用したい値を入れる構造体を作成
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

// グローバルスコープで変数定義をしておくことによりmain関数スコープ内でも変数にアクセスできる様にしているのではないだろうか
var Config ConfigList

// example func xxx (引数) (戻り値の型){}
/*
	func swap(x, y string) (string, string) {
		return x, y
	}
*/

/*
	複数の変数名を一気に定義することが可能

	下記コードのように戻り値に名前をつけることが可能
	この様にすることにより関数の最初に定義した変数名として使用することが可能になる
	func split(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}
	func xx() {
		fmt.Println(split(17))
	}
*/

// 1番先に呼ばれる関数になる
func init() {
	// fileを読み込む
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		// Sectionは[〜]〜を記述する 環境変数名をkeyにintなのでMustIntを選択
		Port:   cfg.Section("web").Key("port").MustInt(),
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
		// String()もし値がなければ空が入る様になっている
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Println("最も早く実行される")
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
}
