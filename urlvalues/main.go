package urlvalues

import (
	"net/url"
	"fmt"
)

func UrlTest()  {
	m := url.Values{"lang":{"en", "zh"}}
	m.Add("item","1")
	m.Add("item","2")
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])
	fmt.Println(m["lang"])
}
