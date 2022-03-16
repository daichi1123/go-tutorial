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
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
