package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var pattern *regexp.Regexp // 検索パターン
var files = []string{}     // 収集したファイル

func main() {
	// コマンドライン引数を得る
	if len(os.Args) <= 2 {
		fmt.Println("[USAGE] collect (pattern) (target)")
		return
	}
	patternStr := os.Args[1]
	target := os.Args[2]
	curDir, _ := os.Getwd()

	// ファイル一覧を得る
	pattern = regexp.MustCompile(patternStr)
	searchFiles(target)

	// コピーを行う
	for _, path := range files {
		fmt.Println("Copy: " + path)
		baseName := filepath.Base(path)
		savePath := filepath.Join(curDir, baseName)
		copyFile(path, savePath)

	}
}

// 指定ディレクトリのファイル一覧を検索
func searchFiles(target string) {
	fs, err := os.ReadDir(target)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fs {
		filePath := filepath.Join(target, file.Name())
		if file.IsDir() {
			// ディレクトリなら再帰的に検索
			searchFiles(filePath)
			continue
		}
		// パターンにマッチするか？
		if pattern.MatchString(file.Name()) {
			files = append(files, filePath)
		}
	}
}

// ファイルをコピーする
func copyFile(inFile, outFile string) bool {
	// ファイルを一度に読み込む
	originalFile, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// ファイルにデータを書き込む
	copyFile, err := os.Create(outFile)
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(copyFile, originalFile)

	return true
}
