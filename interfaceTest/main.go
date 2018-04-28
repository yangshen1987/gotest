package interfaceTest

import (
	"strings"
	"fmt"
)

func NewReaders(p []byte)(n int, err error)  {

	var tmp []byte
	str := string(p[:])
	info := strings.NewReader(str)
	res, err := info.Read(tmp)
	fmt.Println(info)
	fmt.Println(tmp)
	return res, err
}

