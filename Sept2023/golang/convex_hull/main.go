package main

import (
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/convex_hull/hull"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/convex_hull/point"
	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/convex_hull/render"
)

func main() {
	points := []point.Point{
		point.NewPointInCartesian(35, 0),
		point.NewPointInCartesian(-40, 25),
		point.NewPointInCartesian(50, 50),
		point.NewPointInCartesian(-20, 60),
		point.NewPointInCartesian(-40, 50),
		point.NewPointInCartesian(35, -15),
		point.NewPointInCartesian(-10, 10),
		point.NewPointInCartesian(-30, 48.5),
		point.NewPointInCartesian(15, 0),
		point.NewPointInCartesian(5, 10),
		point.NewPointInCartesian(-50, 10),
		point.NewPointInCartesian(5, 40),
		point.NewPointInCartesian(-35, -15),
		point.NewPointInCartesian(25, -10),
		point.NewPointInCartesian(37.5, 20),
		point.NewPointInCartesian(30, -5),
		point.NewPointInCartesian(30, 40),
		point.NewPointInCartesian(-35, -35),
		point.NewPointInCartesian(20, -25),
		point.NewPointInCartesian(-10, -30),
		point.NewPointInCartesian(50, -40),
	}

	connectedPoints := hull.GetConvexHull(points)

	render.PlotPoints(points, connectedPoints)
}
