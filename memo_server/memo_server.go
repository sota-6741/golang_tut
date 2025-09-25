package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

const saveFile = "memo.text" // データファイルの保存先

func main() {
	// ハンドラの登録
	http.HandleFunc("/", readMemoHandler)
	http.HandleFunc("/w", writeMemoHandler)
	// サーバーの起動
	fmt.Print("memo server - [URL] http://localhost:8888/\n")
	http.ListenAndServe(":8888", nil)
}

// ルートへアクセスしたときメモを表示
func readMemoHandler(w http.ResponseWriter, r *http.Request) {
	// データファイルを開く
	text, err := os.ReadFile(saveFile)
	if err != nil {
		text = []byte("ここにメモを記入してください．")
	}
	// HTMLのフォームを返す
	htmlText := html.EscapeString(string(text))
	s := "<html>" +
		"<style>textarea {width:99%; height:200px;}</style>" +
		"<form method='POST' action='/w'>" +
		"<textarea name='text'>" + htmlText + "</textarea>" +
		"<input type='submit' value='保存' /></form></html>"
	w.Write([]byte(s))
}

func writeMemoHandler(w http.ResponseWriter, r *http.Request) {
	// 投稿されたフォームの解析
	r.ParseForm()
	if len(r.Form["text"]) == 0 {
		w.Write([]byte("フォームから投稿してください．"))
		return
	}
	text := r.Form["text"][0]
	// データファイルへ書き込む
	os.WriteFile(saveFile, []byte(text), 0644)
	fmt.Println("save:", text)
	// ルートページへリダイレクトして戻る
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
