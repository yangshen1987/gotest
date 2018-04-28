package main

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "omt trailing new linx")
var sep = flag.String("sep", " ", "separator")

func main()  {
	 x := "hello"
	 for i:=0;i<len(x); i++{
	 	x := x[i]
	 	if x != '!'{
	 		x := x + 'A' - 'a'
	 		fmt.Printf("%c", x)
		}
	 }
	}

