package main

import (
	"fmt"
	"test/geometry"
	"test/intList"
	"test/urlvalues"
	"test/coloredpoint"
	"image/color"
)


func main()  {
	point := []geometry.Point{{1,1},{5,1},{5,4},{1,1}}
	sum := geometry.SumPoint(point)
	fmt.Println(sum)
	point2 := geometry.Path{{1,1},{5,1},{5,4},{1,1}}
	sum2 := geometry.SumPoint2(point2)
	fmt.Println(sum2)
	r := &geometry.Point{1,2}
	r.ScaleBy(2)
	fmt.Println(*r)
	p := geometry.Point{2,3}
	p.ScaleBy(2)
	fmt.Println(p)
	s := r.Distance(p)
	fmt.Println(s)
	st2 := &intList.IntList{}
	sttt := &intList.IntList{4, st2}
	struct1 := intList.IntList{3, sttt}
	s1 := struct1.Sum()
	fmt.Println(s1)
	s2 := sttt.Sum()
	fmt.Println(s2)
	urlvalues.UrlTest()
	red := color.RGBA{255,0,0,255}
	var tp = coloredpoint.ColoredPoint{coloredpoint.Point{1,2}, red}
	tp.ScaleBy(2)
	fmt.Println(tp)
	sp := geometry.Point{1,2}
	sq := geometry.Point{4,6}
	distanceFromP := sp.Distance
	fmt.Println(distanceFromP(sq))

}
