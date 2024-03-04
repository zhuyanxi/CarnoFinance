package helper

import (
	"math"
)

// func RSRS(list []float64) float64 {
// 	// least square method
// 	return 0
// }

func RSRS(closePrice []float64) float64 {
	// var df []float64
	// for _, v := range closePrice {
	// 	df = append(df, v)
	// }
	df := closePrice

	var y []float64
	for _, v := range df {
		y = append(y, math.Log(v))
	}

	var x []float64
	for i := 0; i < len(df); i++ {
		x = append(x, float64(i))
	}

	var sumX, sumY, sumXY, sumXSquare float64
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
	}

	n := float64(len(x))
	slope := (n*sumXY - sumX*sumY) / (n*sumXSquare - sumX*sumX)
	intercept := (sumY - slope*sumX) / n

	annualizedReturns := math.Pow(math.Exp(slope), 250) - 1

	var sumError float64
	for i := 0; i < len(y); i++ {
		sumError += math.Pow(y[i]-(slope*x[i]+intercept), 2)
	}
	rSquared := 1 - (sumError / ((n - 1) * Variance(y, 1)))

	score := annualizedReturns * rSquared
	return score
}

/*
*
参数：接受一个[]float64类型的数据切片以及一个整数ddof，表示自由度（degrees of freedom）。
返回值：返回一个float64类型的值，表示数据的方差。
功能：方差是一组数据与其均值之差的平方和的均值，它表示数据的离散程度。
mean 是数据集的均值
n 是数据点的数量
ddof 是自由度

注意：方差通常用于衡量数据的分散程度，ddof参数用于指定计算方差时的自由度，一般情况下为1。
*/
func Variance(data []float64, ddof int) float64 {
	n := float64(len(data))
	mean := Mean(data)
	sumSquaredDiff := 0.0
	for _, x := range data {
		sumSquaredDiff += (x - mean) * (x - mean)
	}
	return sumSquaredDiff / (n - float64(ddof))
}

func Mean(data []float64) float64 {
	var sum float64
	for _, x := range data {
		sum += x
	}
	return sum / float64(len(data))
}
