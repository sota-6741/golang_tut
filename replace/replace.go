package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// HTMLファイル一覧を得る
	files, err := filepath.Glob("*.html")
	if err != nil {
		fmt.Printf("HTMLファイルの取得に失敗しました: %v\n", err)
	}
	// ファイルを一つづつ処理
	for _, fileName := range files {
		fileReplace(fileName, "hoge@example.com", "fuga@example.com")
	}
}

// ファイルの中のsrcをdstに置換する関数
func fileReplace(fileName, src, dst string) {
	// ファイルを一括で読む
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("ファイルの読み込みに失敗しました: %v\n", err)
	}

	// 文字列置換
	lines := strings.Replace(string(bytes), src, dst, -1)

	// ファイルへ書き込む
	result := []byte(lines)
	resultFileName := makeResultFilename(fileName)
	os.WriteFile(resultFileName, []byte(result), 0666)
	fmt.Println("ok:", resultFileName)
}

// 出力ファイル名の生成
func makeResultFilename(fileName string) string {
	ext := filepath.Ext(fileName)
	name := strings.TrimSuffix(fileName, ext)
	outputFileName := "result_" + name + ext

	return outputFileName
}
