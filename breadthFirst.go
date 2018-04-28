package main

import (
	"test/findlink3"
	"os"
)

func main()  {
	findlink3.BreadtgFirst(findlink3.Crawl, os.Args[1:])
}
