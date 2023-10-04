package render

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/convex_hull/point"
)

func PlotPoints(points, connectedPoints []point.Point) {
	width := 1000
	height := 1000

	graph := gg.NewContext(width, height)
	graph.SetColor(color.White)
	graph.Clear()
	graph.Translate(500, -500)
	graph.InvertY()
	graph.Scale(5, 5)
	for _, point := range points {
		graph.DrawPoint(point.GetX(), point.GetY(), 5)
	}

	graph.SetRGB(1, 0, 0)
	graph.Fill()
	for i, point := range connectedPoints {
		if i == 0 {
			graph.MoveTo(point.GetX(), point.GetY())
		} else if i < len(connectedPoints)-1 {
			graph.LineTo(point.GetX(), point.GetY())
		} else {
			graph.LineTo(point.GetX(), point.GetY())
			graph.LineTo(connectedPoints[0].GetX(), connectedPoints[0].GetY())
		}
	}
	graph.SetRGB(0, 0, 1)
	graph.Stroke()
	graph.SavePNG("show.png")
}
