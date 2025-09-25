package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// カレントディレクトリは以下のファイル一覧を得て出力
	for _, file := range GetFiles(".") {
		fmt.Println(file)
	}
}

// 再起的にファイルの一覧を得る
func GetFiles(directory string) []string {
	var result_files []string
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filepath := filepath.Join(directory, file.Name())
		if file.IsDir() {
			result_files = append(result_files, GetFiles(filepath)...)
			continue
		}
		result_files = append(result_files, filepath)
	}
	return result_files
}
