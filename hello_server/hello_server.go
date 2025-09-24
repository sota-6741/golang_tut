package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler) // ハンドラの登録
	http.ListenAndServe(":8888", nil)  // サーバーを起動
}

// HelloHandler サーバーの処理内容を記述
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
