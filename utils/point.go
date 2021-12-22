package utils

type Point struct {
	X int
	Y int
}

var North = Point{0, 1}
var South = Point{0, -1}
var East = Point{1, 0}
var West = Point{-1, 0}

var Directions = []Point{North, South, East, West}


func (p Point) Add(p2 Point) Point {
	return Point{p.X + p2.X, p.Y + p2.Y}
}

func (p Point) Scale(factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func (p Point) Right() Point {
	return Point{p.Y, -p.X}
}

func (p Point) Left() Point {
	return Point{-p.Y, p.X}
}