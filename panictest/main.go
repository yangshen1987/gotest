package panictest

import (
	"fmt"
	"runtime"
	"os"
)

func F(x int)  {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d/n",x)
	F(x -1)
}
func PrintStrack()  {
	var buf  [4096]byte
	n:= runtime.Stack(buf[:], false)
	os.Stdin.Write(buf[:n])
}
func panicTest(x int)  {
	if x == 0 {
		panic("this panic")
	}
}
func Testrecover()(err error)  {
	defer func() {
		if p := recover(); p!= nil{
			err = fmt.Errorf("interl error :%v", p)
		}
	}()
	panicTest(0)
	return  err
}
