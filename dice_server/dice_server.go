package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	// サーバーの起動
	http.HandleFunc("/", DiceHandler)
	http.ListenAndServe(":8888", nil)
}

// DiceHandler サーバーの処理内容を記述
func DiceHandler(w http.ResponseWriter, r *http.Request) {
	// サイコロの目と出力メッセージを生成
	diceRoll := rand.Intn(6) + 1
	message := fmt.Sprintf("サイコロの目は，%d", diceRoll)
	// 書き出す
	w.Write([]byte(message))
}
