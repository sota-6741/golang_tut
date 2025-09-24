package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// ファイル一覧を得る
	files, err := filepath.Glob("*")
	if err != nil {
		fmt.Printf("ファイル読み込みに失敗しました: %v\n", err)
	}
	// ファイルを一つずつ表示する
	for i, name := range files {
		fmt.Println(i, "=", name)
	}
}
