package kalmanfilter

import (
	"math"
)

// RSRS 计算带有卡尔曼滤波器的 RSRS 动量得分 (基于 t 统计量)
// 参数 closePrice: 历史收盘价切片 (建议长度为 25 或以上)
// 返回值: 经风险调整后的动态斜率得分 (t 统计量)
func RSRS(closePrice []float64) float64 {
	n := len(closePrice)
	if n < 3 {
		return 0.0
	}

	// 1. 将价格转换为对数价格
	y := make([]float64, n)
	for i, v := range closePrice {
		if v <= 0 {
			v = 1e-5 // 避免 log(0) 导致 NaN
		}
		y[i] = math.Log(v)
	}

	// 2. 自适应计算观测噪声 R
	// 计算滚动对数收益率的方差作为波动率度量
	logReturns := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		logReturns[i] = y[i+1] - y[i]
	}
	realizedVol := StdDev(logReturns, 1)

	// rMultiplier 为波动率转换系数，通常设为 0.5
	rMultiplier := 0.5
	rAdaptive := realizedVol * realizedVol * rMultiplier
	if rAdaptive < 1e-5 {
		rAdaptive = 1e-5 // 设定下限，防止系统过于灵敏
	}

	// 3. OLS 线性回归初始化状态 (Warm-up)
	// 构造自变量 x = [0, 1, 2, ... n-1]
	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = float64(i)
	}

	var sumX, sumY, sumXY, sumXSquare float64
	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSquare += x[i] * x[i]
	}
	num := float64(n)
	slopeInit := (num*sumXY - sumX*sumY) / (num*sumXSquare - sumX*sumX)
	interceptInit := (sumY - slopeInit*sumX) / num

	// 4. 运行二维卡尔曼滤波器
	// 状态向量 theta = [slope, intercept]^T
	slope := slopeInit
	intercept := interceptInit

	// 协方差矩阵 P (对称矩阵)
	// P00: slope 的方差, P11: intercept 的方差, P01: 协方差
	p00, p11, p01 := 1.0, 1.0, 0.0

	// 过程噪声 Q_noise (系统状态随机游走的标准差)
	qNoise := 1e-4

	for t := 0; t < n; t++ {
		// 当前时间步的自变量（观测矩阵 H = [x_val, 1.0]）
		xVal := x[t]
		yObs := y[t]

		// 【预测步骤 (Time Update)】
		// theta_pred = theta (因为状态转移矩阵为单位矩阵 I)
		// P_pred = P + Q
		pPred00 := p00 + qNoise
		pPred11 := p11 + qNoise
		pPred01 := p01

		// 【更新步骤 (Measurement Update)】
		// 1. 计算创新协方差 S = H * P_pred * H^T + R
		s := xVal*(xVal*pPred00+pPred01) + (xVal*pPred01 + pPred11) + rAdaptive

		// 2. 计算卡尔曼增益 K = P_pred * H^T / S
		k0 := (xVal*pPred00 + pPred01) / s
		k1 := (xVal*pPred01 + pPred11) / s

		// 3. 计算残差 (Measurement Residual)
		residual := yObs - (xVal*slope + intercept)

		// 4. 更新状态估计
		slope += k0 * residual
		intercept += k1 * residual

		// 5. 更新协方差矩阵 P = (I - K * H) * P_pred
		pNew00 := (1.0-xVal*k0)*pPred00 - k0*pPred01
		pNew01 := (1.0-xVal*k0)*pPred01 - k0*pPred11
		pNew11 := -xVal*k1*pPred01 + (1.0-k1)*pPred11

		p00 = pNew00
		p01 = pNew01
		p11 = pNew11
	}

	// 5. 计算内生 t 统计量
	// 最终提炼出的斜率估计误差标准差
	slopeStdErr := math.Sqrt(p00)
	if slopeStdErr <= 0 {
		slopeStdErr = 1e-8
	}

	// t 统计量 = 斜率 / 标准误
	tStat := slope / slopeStdErr

	return tStat
}

// Variance 计算一维切片的方差
func Variance(data []float64, ddof int) float64 {
	n := float64(len(data))
	if n <= float64(ddof) {
		return 0.0
	}
	mean := Mean(data)
	sumSquaredDiff := 0.0
	for _, x := range data {
		sumSquaredDiff += (x - mean) * (x - mean)
	}
	return sumSquaredDiff / (n - float64(ddof))
}

// StdDev 计算一维切片的标准差
func StdDev(data []float64, ddof int) float64 {
	return math.Sqrt(Variance(data, ddof))
}

// Mean 计算一维切片的算术平均值
func Mean(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}
	var sum float64
	for _, x := range data {
		sum += x
	}
	return sum / float64(len(data))
}
