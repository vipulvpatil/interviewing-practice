package point

import (
	"math"
)

type Point interface {
	GetX() float64
	GetY() float64
	GetAngle() float64
	GetLength() float64
}

type point struct {
	x      float64
	y      float64
	angle  float64
	length float64
}

func (p *point) GetX() float64 {
	return p.x
}

func (p *point) GetY() float64 {
	return p.y
}

func (p *point) GetAngle() float64 {
	return p.angle
}

func (p *point) GetLength() float64 {
	return p.length
}

func NewPointInCartesian(x, y float64) Point {
	p := point{
		x:      x,
		y:      y,
		angle:  getAngleFromCartesian(x, y),
		length: getLengthFromCartesian(x, y),
	}
	return &p
}

func NewPointInPolar(length, angle float64) Point {
	p := point{
		x:      getXFromPolar(length, angle),
		y:      getYFromPolar(length, angle),
		angle:  normalizeAngle(angle),
		length: length,
	}
	return &p
}

// normalizeAngle normalizes a given angle value to be between 0 to 359
func normalizeAngle(angle float64) float64 {
	for angle >= 360 {
		angle -= 360
	}
	for angle <= -360 {
		angle += 360
	}
	if angle < 0 {
		angle += 360
	}
	return angle
}

func getXFromPolar(length, angle float64) float64 {
	return length * math.Cos(angle)
}

func getYFromPolar(length, angle float64) float64 {
	return length * math.Sin(angle)
}

func getAngleFromCartesian(x, y float64) float64 {
	if x == 0 {
		if y == 0 {
			return 0
		}
		return 90 * math.Abs(y)
	}
	angleInRadian := math.Atan(y / x)
	return normalizeAngle(angleInRadian * 180 / math.Pi)
}

func getLengthFromCartesian(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
