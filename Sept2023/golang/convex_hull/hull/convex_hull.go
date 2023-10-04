package hull

import (
	"math"

	"github.com/vipulvpatil/interviewing-practice/Sept2023/golang/convex_hull/point"
	"golang.org/x/exp/slices"
)

func GetConvexHull(points []point.Point) []point.Point {
	if len(points) < 3 {
		return points
	}
	sortByY(points)
	lowestPoint := points[0]
	sortByAngleFromGivenPoint(points[1:], lowestPoint)
	hullPoints := []point.Point{}
	hullPoints = append(hullPoints, points[0])
	hullPoints = append(hullPoints, points[1])
	for i := 2; i < len(points); {
		hullPointsLen := len(hullPoints)
		aPoint := hullPoints[hullPointsLen-2]
		bPoint := hullPoints[hullPointsLen-1]
		cPoint := points[i]
		if !isCCWTurn(aPoint, bPoint, cPoint) {
			hullPoints = hullPoints[:hullPointsLen-1]
		} else {
			hullPoints = append(hullPoints, cPoint)
			i++
		}
	}
	return hullPoints
}

func turn(a, b, c point.Point) float64 {
	return a.GetX()*(b.GetY()-c.GetY()) + b.GetX()*(c.GetY()-a.GetY()) + c.GetX()*(a.GetY()-b.GetY())
}

func isCCWTurn(a, b, c point.Point) bool {
	return turn(a, b, c) > 0
}

func sortByY(points []point.Point) {
	slices.SortFunc[[]point.Point](points, func(a point.Point, b point.Point) int {
		if a.GetY() < b.GetY() {
			return -1
		}
		if a.GetY() > b.GetY() {
			return 1
		}
		if a.GetX() > b.GetX() {
			return -1
		}
		if a.GetX() < b.GetX() {
			return 1
		}
		return 0
	})
}

func sortByAngleFromGivenPoint(points []point.Point, givenPoint point.Point) {
	slices.SortFunc[[]point.Point](points, func(a point.Point, b point.Point) int {
		turnDirection := turn(givenPoint, a, b)
		if turnDirection == 0 {
			distASqr := math.Pow(a.GetX()-givenPoint.GetX(), 2) + math.Pow(a.GetY()-givenPoint.GetY(), 2)
			distBSqr := math.Pow(b.GetX()-givenPoint.GetX(), 2) + math.Pow(b.GetY()-givenPoint.GetY(), 2)
			if distASqr < distBSqr {
				return -1
			} else if distASqr > distBSqr {
				return 1
			}
			return 0
		}

		if turnDirection > 0 {
			return -1
		}

		return 1
	})
}
