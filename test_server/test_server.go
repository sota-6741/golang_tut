package main

import "net/http"

func main() {
	// URIに対応するハンドラを登録
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/aaa", aaaHandler)
	http.HandleFunc("/bbb", bbbHandler)
	// サーバーを起動
	http.ListenAndServe(":8888", nil)
}

// ハンドラを定義
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello")) // Helloと出力
}

func aaaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aaa")) // aaaと出力
}

func bbbHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bbb")) // bbbと出力
}
