package main

// ライブラリの取り込み
import (
	"fmt"
	"math"
)

// 体重(kg)と身長(cm)を指定
const weight = 95
const height = 181

func main() {
	// BMIの計算
	hm := height / 100.0
	bmi := weight / math.Pow(hm, 2)
	bestWeight := math.Pow(hm, 2) * 22.0
	per := weight / bestWeight * 100

	// 結果の表示
	fmt.Printf("BMI=%f, 肥満度=%.0f\n", bmi, per)
}
