package main

import (
	"log"

	"image"

	"image/png"
	"os"

	"image/color"

	"time"

	"math"

	"github.com/ZzEeKkAa/fuzzy-logic/fuzzy"
	"github.com/golang/geo/r2"
	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	var s Surface
	s.SetCore(func(x, y float64) float64 {
		//return 0
		return math.Sin((x - 10) * (y - 5) / 5)
	})
	//for i := 0.; i < 10; i += 0.01 {
	//	log.Println(s.GetVehicleAngels(i, 2*i, 0, 0.05, 0.02))
	//}

	deviationSet := fuzzy.Set{
		{r2.Point{X: -180, Y: 1}, r2.Point{X: -180, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -50, Y: 0}},
		{r2.Point{X: -100, Y: 0}, r2.Point{X: -40, Y: 1}, r2.Point{X: -35, Y: 1}, r2.Point{X: -15, Y: 0}},
		{r2.Point{X: -20, Y: 0}, r2.Point{X: -5, Y: 1}, r2.Point{X: 5, Y: 1}, r2.Point{X: 20, Y: 0}},
		{r2.Point{X: 15, Y: 0}, r2.Point{X: 35, Y: 1}, r2.Point{X: 40, Y: 1}, r2.Point{X: 60, Y: 0}},
		{r2.Point{X: 50, Y: 0}, r2.Point{X: 100, Y: 1}, r2.Point{X: 180, Y: 1}, r2.Point{X: 180, Y: 1}},
	}
	tiltSet := fuzzy.Set{
		{r2.Point{X: -90, Y: 1}, r2.Point{X: -90, Y: 1}, r2.Point{X: -30, Y: 1}, r2.Point{X: -25, Y: 0}},
		{r2.Point{X: -30, Y: 0}, r2.Point{X: -25, Y: 1}, r2.Point{X: -10, Y: 1}, r2.Point{X: -5, Y: 0}},
		{r2.Point{X: -10, Y: 0}, r2.Point{X: -5, Y: 1}, r2.Point{X: 5, Y: 1}, r2.Point{X: 10, Y: 0}},
		{r2.Point{X: 5, Y: 0}, r2.Point{X: 10, Y: 1}, r2.Point{X: 25, Y: 1}, r2.Point{X: 30, Y: 0}},
		{r2.Point{X: 25, Y: 0}, r2.Point{X: 30, Y: 1}, r2.Point{X: 90, Y: 1}, r2.Point{X: 90, Y: 1}},
	}

	turnSet := fuzzy.Set{
		{r2.Point{X: -60, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -40, Y: 1}, r2.Point{X: -30, Y: 0}},
		{r2.Point{X: -40, Y: 0}, r2.Point{X: -30, Y: 1}, r2.Point{X: -10, Y: 1}, r2.Point{X: -5, Y: 0}},
		{r2.Point{X: -15, Y: 0}, r2.Point{X: 0, Y: 1}, r2.Point{X: 0, Y: 1}, r2.Point{X: 15, Y: 0}},
		{r2.Point{X: 5, Y: 0}, r2.Point{X: 10, Y: 1}, r2.Point{X: 30, Y: 1}, r2.Point{X: 40, Y: 0}},
		{r2.Point{X: 30, Y: 0}, r2.Point{X: 40, Y: 1}, r2.Point{X: 60, Y: 1}, r2.Point{X: 60, Y: 1}},
	}

	am := fuzzy.AssociativeMemory{}
	am.Init(deviationSet, tiltSet)
	am.Set(
		&turnSet[0], &turnSet[0], &turnSet[1], &turnSet[2], &turnSet[2],
		&turnSet[0], &turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3],
		&turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3], &turnSet[3],
		&turnSet[1], &turnSet[1], &turnSet[3], &turnSet[3], &turnSet[4],
		&turnSet[1], &turnSet[2], &turnSet[3], &turnSet[4], &turnSet[4],
	)

	traectory := BuildTraectory(0.2, 0.1, r2.Point{X: 0, Y: 0}, 0, r2.Point{X: 20, Y: 10}, 0.1, am, s)
	log.Println(traectory)

	k := 20.
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	drawSurface(img, s, k, 2, -2)
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(color.RGBA{})
	gc.SetLineWidth(1)

	//gc.SetStrokeColor(color.RGBA{G: 255, A: 255})
	gc.MoveTo(traectory[0].X, traectory[0].Y)
	for _, p := range traectory {
		gc.LineTo(p.X*k, p.Y*k)
		gc.MoveTo(p.X*k, p.Y*k)
	}
	gc.Close()
	gc.FillStroke()
	f, _ := os.Create("path_" + time.Now().String() + ".png")
	png.Encode(f, img)
	f.Close()
}

func drawSurface(img *image.RGBA, s Surface, k float64, max, min float64) {
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			//log.Println(i, j)
			h := s.GetHeight(float64(i)/k, float64(j)/k)
			k := 0.
			if h > max {
				k = 1.
			} else if h > min {
				k = (h - min) / (max - min)
			}
			img.SetRGBA(i, j, color.RGBA{R: uint8(255 * k), G: uint8(255 * (1 - k)), B: 255, A: 255})
		}
	}
}
