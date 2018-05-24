package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"time"

	"github.com/ZzEeKkAa/fuzzy-logic/fuzzy"
	"github.com/golang/geo/r2"
	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	var surface Surface
	h := func(x, y float64) float64 {
		//return 0
		//return math.Sin((x - 10) * (y - 5) / 5)
		return math.Sin((x) * (y) / 5)
	}
	surface.SetCore(h)

	//deviationSet := fuzzy.Set{
	//	{r2.Point{X: -180, Y: 1}, r2.Point{X: -180, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -50, Y: 0}},
	//	{r2.Point{X: -100, Y: 0}, r2.Point{X: -40, Y: 1}, r2.Point{X: -35, Y: 1}, r2.Point{X: -15, Y: 0}},
	//	{r2.Point{X: -20, Y: 0}, r2.Point{X: -5, Y: 1}, r2.Point{X: 5, Y: 1}, r2.Point{X: 20, Y: 0}},
	//	{r2.Point{X: 15, Y: 0}, r2.Point{X: 35, Y: 1}, r2.Point{X: 40, Y: 1}, r2.Point{X: 60, Y: 0}},
	//	{r2.Point{X: 50, Y: 0}, r2.Point{X: 100, Y: 1}, r2.Point{X: 180, Y: 1}, r2.Point{X: 180, Y: 1}},
	//}
	//tiltSet := fuzzy.Set{
	//	{r2.Point{X: -90, Y: 1}, r2.Point{X: -90, Y: 1}, r2.Point{X: -30, Y: 1}, r2.Point{X: -25, Y: 0}},
	//	{r2.Point{X: -30, Y: 0}, r2.Point{X: -25, Y: 1}, r2.Point{X: -10, Y: 1}, r2.Point{X: -5, Y: 0}},
	//	{r2.Point{X: -10, Y: 0}, r2.Point{X: -5, Y: 1}, r2.Point{X: 5, Y: 1}, r2.Point{X: 10, Y: 0}},
	//	{r2.Point{X: 5, Y: 0}, r2.Point{X: 10, Y: 1}, r2.Point{X: 25, Y: 1}, r2.Point{X: 30, Y: 0}},
	//	{r2.Point{X: 25, Y: 0}, r2.Point{X: 30, Y: 1}, r2.Point{X: 90, Y: 1}, r2.Point{X: 90, Y: 1}},
	//}
	//turnSet := fuzzy.Set{
	//	{r2.Point{X: -60, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -40, Y: 1}, r2.Point{X: -30, Y: 0}},
	//	{r2.Point{X: -40, Y: 0}, r2.Point{X: -30, Y: 1}, r2.Point{X: -10, Y: 1}, r2.Point{X: -5, Y: 0}},
	//	{r2.Point{X: -15, Y: 0}, r2.Point{X: 0, Y: 1}, r2.Point{X: 0, Y: 1}, r2.Point{X: 15, Y: 0}},
	//	{r2.Point{X: 5, Y: 0}, r2.Point{X: 10, Y: 1}, r2.Point{X: 30, Y: 1}, r2.Point{X: 40, Y: 0}},
	//	{r2.Point{X: 30, Y: 0}, r2.Point{X: 40, Y: 1}, r2.Point{X: 60, Y: 1}, r2.Point{X: 60, Y: 1}},
	//}

	deviationSet := fuzzy.Set{
		{r2.Point{X: -180, Y: 1}, r2.Point{X: -180, Y: 1}, r2.Point{X: -70, Y: 1}, r2.Point{X: -30, Y: 0}},
		{r2.Point{X: -100, Y: 0}, r2.Point{X: -40, Y: 1}, r2.Point{X: -35, Y: 1}, r2.Point{X: -15, Y: 0}},
		{r2.Point{X: -20, Y: 0}, r2.Point{X: -5, Y: 1}, r2.Point{X: 5, Y: 1}, r2.Point{X: 20, Y: 0}},
		{r2.Point{X: 15, Y: 0}, r2.Point{X: 35, Y: 1}, r2.Point{X: 40, Y: 1}, r2.Point{X: 100, Y: 0}},
		{r2.Point{X: 30, Y: 0}, r2.Point{X: 70, Y: 1}, r2.Point{X: 180, Y: 1}, r2.Point{X: 180, Y: 1}},
	}
	tiltSet := fuzzy.Set{
		{r2.Point{X: -90, Y: 1}, r2.Point{X: -90, Y: 1}, r2.Point{X: -50, Y: 1}, r2.Point{X: -25, Y: 0}},
		{r2.Point{X: -30, Y: 0}, r2.Point{X: -20, Y: 1}, r2.Point{X: -15, Y: 1}, r2.Point{X: -5, Y: 0}},
		{r2.Point{X: -10, Y: 0}, r2.Point{X: -2, Y: 1}, r2.Point{X: 2, Y: 1}, r2.Point{X: 10, Y: 0}},
		{r2.Point{X: 5, Y: 0}, r2.Point{X: 15, Y: 1}, r2.Point{X: 20, Y: 1}, r2.Point{X: 30, Y: 0}},
		{r2.Point{X: 25, Y: 0}, r2.Point{X: 50, Y: 1}, r2.Point{X: 90, Y: 1}, r2.Point{X: 90, Y: 1}},
	}
	turnSet := fuzzy.Set{
		{r2.Point{X: -60, Y: 1}, r2.Point{X: -60, Y: 1}, r2.Point{X: -40, Y: 1}, r2.Point{X: -25, Y: 0}},
		{r2.Point{X: -40, Y: 0}, r2.Point{X: -25, Y: 1}, r2.Point{X: -20, Y: 1}, r2.Point{X: -5, Y: 0}},
		{r2.Point{X: -15, Y: 0}, r2.Point{X: 0, Y: 1}, r2.Point{X: 0, Y: 1}, r2.Point{X: 15, Y: 0}},
		{r2.Point{X: 5, Y: 0}, r2.Point{X: 20, Y: 1}, r2.Point{X: 25, Y: 1}, r2.Point{X: 40, Y: 0}},
		{r2.Point{X: 25, Y: 0}, r2.Point{X: 40, Y: 1}, r2.Point{X: 60, Y: 1}, r2.Point{X: 60, Y: 1}},
	}

	amDown := fuzzy.AssociativeMemory{}
	amDown.Init(deviationSet, tiltSet)
	amDown.Set(
		&turnSet[0], &turnSet[0], &turnSet[1], &turnSet[2], &turnSet[2],
		&turnSet[0], &turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3],
		&turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3], &turnSet[3],
		&turnSet[1], &turnSet[1], &turnSet[3], &turnSet[3], &turnSet[4],
		&turnSet[1], &turnSet[2], &turnSet[3], &turnSet[4], &turnSet[4],
	)

	amUp := fuzzy.AssociativeMemory{}
	amUp.Init(deviationSet, tiltSet)
	amUp.Set(
		&turnSet[1], &turnSet[2], &turnSet[3], &turnSet[4], &turnSet[4],
		&turnSet[1], &turnSet[1], &turnSet[3], &turnSet[3], &turnSet[4],
		&turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3], &turnSet[3],
		&turnSet[0], &turnSet[1], &turnSet[1], &turnSet[2], &turnSet[3],
		&turnSet[0], &turnSet[0], &turnSet[1], &turnSet[2], &turnSet[2],
	)

	var vehicle = Vehicle{a: 0.1, b: 0.2, x: 0, y: 0, alpha: 0}
	var endPoint = r2.Point{X: 20, Y: 10}

	trajectory := BuildTrajectoryNew(vehicle, endPoint, 0.01, 0.1, amDown, amDown, surface.GetHeight)
	//log.Println(trajectory)

	score := ScoreTraectory(vehicle, trajectory, surface)

	log.Println("Score: ", score)

	k := 20.
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	drawSurface(img, surface, k, 2, -2)

	lastPoint := trajectory[0]
	for _, p := range trajectory {
		//gc.SetStrokeColor(color.RGBA{G: 255, A: 255})
		gc := draw2dimg.NewGraphicContext(img)
		gc.SetStrokeColor(image.Black)
		gc.SetFillColor(color.RGBA{})
		gc.SetLineWidth(1)
		gc.MoveTo(lastPoint.x*k, lastPoint.y*k)

		vehicle.ApplyPosition(&p)
		tang, _ := vehicle.FindDegrees(surface.GetHeight)

		c := uint8((tang + 90) / 180 * 255)
		gc.SetStrokeColor(color.RGBA{R: c, G: c, B: c, A: 255})
		//gc.LineTo(p.X*k, p.Y*k)
		//gc.MoveTo(p.X*k, p.Y*k)
		gc.LineTo(p.x*k, p.y*k)
		//gc.MoveTo(p.x*k, p.y*k)
		gc.Close()
		gc.FillStroke()
		lastPoint = p
	}

	img.Set(int(endPoint.X*k), int(endPoint.Y*k), color.RGBA{255, 0, 0, 255})
	img.Set(int(endPoint.X*k-1), int(endPoint.Y*k), color.RGBA{255, 0, 0, 255})
	img.Set(int(endPoint.X*k+1), int(endPoint.Y*k), color.RGBA{255, 0, 0, 255})
	img.Set(int(endPoint.X*k), int(endPoint.Y*k-1), color.RGBA{255, 0, 0, 255})
	img.Set(int(endPoint.X*k), int(endPoint.Y*k+1), color.RGBA{255, 0, 0, 255})
	f, _ := os.Create("dist/path_" + time.Now().String() + ".png")
	//f, _ := os.Create("dist/path_tmp.png")
	png.Encode(f, img)
	f.Close()
}

func drawSurface(img *image.RGBA, s Surface, k float64, max, min float64) {
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
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
