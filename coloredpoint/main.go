package coloredpoint

import "image/color"

type Point struct {
	X, Y float64
}
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point)ScaleBy(factor float64)  {
	p.Y *= factor
	p.X *= factor
}
func (p Point)Add(q Point) Point  {
	return Point{p.X+q.X, q.Y+p.Y}
}
func (p Point)Sub(q Point) Point  {
	return Point{p.X-q.X, p.Y-q.Y}
}

type Path []Point

func (path Path)TranslateBy(offset Point, add bool)  {
	var op func(p,q Point) Point
	if add {
		op = Point.Add
	}else{
		op = Point.Sub
	}
	for i := range path{
		path[i] = op(path[i], offset)
	}
}

