package geometry

import "math"

type Point struct {
	X,Y float64
}
type Path []Point

func Distance(p,q Point) float64  {
	return math.Hypot(q.X-p.Y, q.Y-p.X)
}
func (p Point)Distance(q Point) float64  {
	return math.Hypot(q.X-p.Y, q.Y-p.X)
}
func (path Path)Distance() float64  {
	 sum := 0.0
	 for i:= range path{
	 	if i>0{
	 		sum += path[i-1].Distance(path[i])
		}
	 }
	 return sum
}
func SumPoint(x []Point) float64  {
	var perim Path
	perim = x
	return perim.Distance()
}
func SumPoint2(x Path) float64 {
	return x.Distance()
}
func (p *Point)ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}
