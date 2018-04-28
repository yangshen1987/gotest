package bytecounter

import (
	"bufio"
)

type ByteCounter int

func (c *ByteCounter)Write(p []byte)(int, error)  {
	*c += ByteCounter(len(p))
	return  len(p), nil
}
func (c *ByteCounter)Count(p []byte)(int, error)  {
	start, _ ,error := bufio.ScanWords(p, true)
	*c += 1
	if start != len(p) && error == nil{
		start, error =c.Count(p[start:])
	}
	return start, error
}
