package calculate

import (
	"math"
)

// 除算 + baseNum桁で表示（baseNum+1桁の値で四捨五入）
func Division(top int, bottom int, baseNum int) float64 {
	bn := math.Pow(10, float64(baseNum))
	raw_ans := float64(top) / float64(bottom)
	ans := math.Round(raw_ans*bn) / bn
	return ans
}

// 除算 + baseNum桁で表示（baseNum+1桁の値で四捨五入）
func DivisionFloat(top float64, bottom float64, baseNum int) float64 {
	bn := math.Pow(10, float64(baseNum))
	raw_ans := float64(top) / float64(bottom)
	ans := math.Round(raw_ans*bn) / bn
	return ans
}

// %表示 + 小数点baseNum桁まで表示（小数点baseNum+1桁の値で四捨五入）
func Rate(top int, bottom int, baseNum int) float64 {
	bn := math.Pow(10, float64(baseNum))
	raw_ans := float64(top) / float64(bottom) * 100
	ans := math.Round(raw_ans*bn) / bn
	return ans
}
