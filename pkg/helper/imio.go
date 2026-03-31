package helper

import (
	"fmt"
)

// MarketData 结构体存储实时行情和技术计算所需的基础数据
type MarketData struct {
	SpotPrice    float64
	FuturePrice  float64
	IV           float64
	Highs        []float64
	Lows         []float64
	DaysToExpiry int
}

// Indicators 存储计算后的 5 重指标结果
type Indicators struct {
	RSRSScore   float64
	RSI6        float64
	RSI12       float64
	BollUpper   float64
	BollMid     float64
	BollLower   float64
	IVRank      float64
	BasisAnnual float64
}

// 模拟数学辅助函数：线性回归斜率 (RSRS 核心)
func CalculateSlope(highs, lows []float64) float64 {
	n := float64(len(highs))
	var sumX, sumY, sumXY, sumXX float64
	for i := 0; i < len(highs); i++ {
		sumX += lows[i]
		sumY += highs[i]
		sumXY += lows[i] * highs[i]
		sumXX += lows[i] * lows[i]
	}
	// 简单线性回归公式: beta = (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	return (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
}

// 策略核心逻辑
func (m *MarketData) ExecuteStrategy() {
	// 1. 计算指标 (此处简化了历史计算过程)
	ind := Indicators{
		RSRSScore:   CalculateSlope(m.Highs, m.Lows), // 实际需配合 Z-Score 标准化
		RSI6:        25.0,                            // 模拟值
		RSI12:       28.0,                            // 模拟值
		BollLower:   m.SpotPrice * 0.98,              // 模拟下轨
		IVRank:      0.85,                            // 模拟高 IV
		BasisAnnual: ((m.SpotPrice - m.FuturePrice) / m.SpotPrice) * (365.0 / float64(m.DaysToExpiry)),
	}

	fmt.Printf("--- 策略监控中 | 基差年化: %.2f%% | RSRS: %.2f ---\n", ind.BasisAnnual*100, ind.RSRSScore)

	// --- 交易触发逻辑 ---

	// 基础底仓：IM 贴水收割 (10% 年化门槛)
	hasLongIM := ind.BasisAnnual > 0.10

	switch {
	// A. 极端超跌共振 (抄底 + 认沽反比例保险)
	case ind.RSRSScore < -1.0 && ind.RSI12 < 30 && m.SpotPrice <= ind.BollLower && ind.IVRank > 0.7:
		fmt.Println("【信号】极端超跌！执行：持有 IM 多头 + 构建 1:3 认沽反比例价差 (MO Put Backspread)")
		// callOrderAPI("MO", "RatioPutBackspread", 3)

	// B. 趋势上行 (右侧追涨 + Gamma 爆发)
	case ind.RSRSScore > 0.7 && ind.RSI12 < 60:
		fmt.Println("【信号】趋势确认！执行：持有 IM 多头 + 构建 1:2 认购反比例价差 (MO Call Backspread)")

	// C. 超买防御 (高位收租)
	case ind.RSI12 > 70 && m.SpotPrice >= ind.BollUpper && ind.IVRank < 0.3:
		fmt.Println("【信号】超买预警！执行：卖出虚值 Call (Delta 0.15) 增厚收益")

	// D. 均值回归或动能消失 (平仓期权)
	case ind.RSRSScore < 0 || (ind.RSI12 > 45 && ind.RSI12 < 55):
		if !hasLongIM {
			fmt.Println("【信号】动能衰竭，清空期权头寸，保留 IM 贴水底仓")
		}
	}
}

func TestIMIO() {
	// 模拟行情数据 (2026年3月环境)
	market := MarketData{
		SpotPrice:    5000.0,
		FuturePrice:  4960.0,
		IV:           0.32,
		DaysToExpiry: 20,
		Highs:        []float64{5100, 5080, 5050}, // 简化数据
		Lows:         []float64{4950, 4920, 4900},
	}

	market.ExecuteStrategy()
}
