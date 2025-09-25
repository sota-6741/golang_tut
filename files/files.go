package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// カレントディレクトリのファイル一覧を得る
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	// 取得した一覧を表示
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
