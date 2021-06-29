package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotLine(title, labelX, labelY, filename string, points map[string]plotter.XYs) {
	rand.Seed(int64(0))

	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = labelX
	p.Y.Label.Text = labelY

	for name, point := range points {
		if err := plotutil.AddLinePoints(p,
			name, point,
		); err != nil {
			panic(err)
		}
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func PlotRandomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

func PlotPoints(dataX, dataY []float64) plotter.XYs {
	if len(dataX) != len(dataY) {
		panic("lens of x and y should be the same")
	}
	pts := make(plotter.XYs, len(dataX))
	for i := range pts {
		pts[i] = plotter.XY{
			X: dataX[i],
			Y: dataY[i],
		}
	}
	return pts
}
