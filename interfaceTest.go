package main

import (
	"test/bytecounter"
	"fmt"
	"test/interfaceTest"
	"io"
)

func main()  {
	var c bytecounter.ByteCounter
	str := []byte("hello yangshen niu")
	c.Count(str)
	fmt.Println(c)
	str1 := []byte("yangshen")
	interfaceTest.NewReaders(str1)

}
func writeString(w io.Writer, s string)(n int, err error)  {
	type stringWriter interface {
		WriteString(string)(n int, err error)
	}
	if sw, ok := w.(stringWriter); ok{
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}
