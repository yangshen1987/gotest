package main

import (
	"fmt"
	"strings"
	"bytes"
	"crypto/sha256"
	"bufio"
	"os"
	"unicode/utf8"
	"io"
	"unicode"
	"encoding/json"
	"log"
	"time"
	"net/url"
	"net/http"
	//"test/toposprt"
	"test/panictest"
)

const IPv4Len  =4
const IssuesURL  = "https://api.github.com/search/issues"
type Weekday int
type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool`json:"color,omitempty"`
	Actors []string
}

var movies  = []Movie{
	{Title:"yangshen", Year:1942, Color:false, Actors:[]string{
		"yangshen yong bu qu fu","shi de  jiushi"}},
	{Title:"caoxiubo", Year:1988,Color:false,Actors:[]string{
		"nihaoa", "wobuhao"}},
}

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thurday
	Friday
	Saturday
)
type Flags uint
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}
type Issue struct {
	Number  int
	HtmlUrl string `json:"html_url"`
	Title   string
	State   string
	User    *User
	CreatedAt time.Time `json:"created_at"`
	Body    string
}
type User struct {
	Login string
	HtmlUrl string `json:"html_url"`
}

const (
	Flagup Flags = 1 <<iota
	FlageBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)
var m = make(map[string]int)
func main()  {

	//const Pi64 float64  = math.Pi
	//var x  float64  = math.Pi
	//var y  = Pi64
	//var z complex128  = complex128(Pi64)
	//var f float64 = 212
	//fmt.Println((f-32)*5/9)
	
	//s := "c/b/a.go"
	//b := "cba/a/2/2/3/2"
	//var r  [IPv4Len]byte
	//fmt.Println(r)
	//fmt.Println(basename(s))
	//fmt.Println(basename2(s))
	//fmt.Println(comma([]string{"1.5","2.5","3.6"}))
	//fmt.Println(checkString(s, b, "/"))
	//fmt.Println(Sunday)
	//fmt.Println(FlageBroadcast)sa
	//fmt.Println(FlagLoopback)
	//fmt.Println(btypeSha256())
	//fmt.Println(popcount.Popcount(25))
	////mapTest()
	//charcount()
	//jsonMake()
	//result ,err :=SearchIssues(os.Args[1:])
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Printf("%d issues :\n", result.TotalCount )
	//for _, item := range result.Items{
	//	fmt.Printf("#5 %9.9s %.55s\n",
	//		item.Number, item.User.Login, item.Title)
	//}
	//order :=toposprt.Topsort(toposprt.Prereqs)
	//for i, coure := range order{
	//	fmt.Printf("%d:\t%s\n", i+1,coure)
	//}
	//fmt.Println("this is end")
	//defer panictest.PrintStrack();
	//panictest.F(3)
	errmsg := panictest.Testrecover()
	fmt.Printf("error msg is %s", errmsg)


}
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}

	}
	for i:=len(s)-1;i>=0 ;i--  {
		if s[i] == '.'{
			s = s[:i]
			break
		}
	}
	return s
}
func basename2(s string) string  {
	slash := strings.LastIndex(s, "/")
	s  = s[slash+1:]
	if dot := strings.LastIndex(s,"."); dot>=0{
		s = s[:dot]
	}
	return s

}
func comma(valuse []string) string  {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range valuse{
		if i>0 {
			buf.WriteString(",")
		}
			fmt.Fprintf(&buf, "%s", v)

	}
	buf.WriteByte(']')
	return buf.String()
}
func checkString(v1 string, v2 string, checkStr string) bool  {
	if strings.LastIndex(v1, checkStr) != strings.LastIndex(v2, checkStr){
		return false
	}
	return true
}
func btypeSha256() string  {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	if c1==c2{
		return "alex"
	}
	return "yangshen"
}
func mapTest()  {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line := input.Text()
		if !seen[line]{
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil{
		fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
		os.Exit(1)
	}
}
func mapKey(list []string) string  {
	return fmt.Sprintf("%q", list)
}
func Add(list []string)  {
	m[mapKey(list)] ++
}
func Count(list []string) int  {
	return m[mapKey(list)]
}
func charcount()  {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax+1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err != nil{
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
		}
		if err == io.EOF{
			break
		}
		if r == unicode.ReplacementChar && n == 1{
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

	}
	fmt.Printf("rune\tcount\n")
	for c,n := range counts{
		fmt.Printf("%d\t%d\n",c,n)
	}
	for i,n := range utflen{
		if i>0{
			fmt.Printf("%d\t%d\n",i,n)
		}
	}
	if invalid >0{
		fmt.Printf("\n%d invalid UTF-8 charachters\n", invalid)
	}
}
func jsonMake()  {
	data ,err := json.MarshalIndent(movies, "", "   ")
	if err != nil{
		log.Fatalf("JSON marshaling failed :%s", err)
	}
	fmt.Printf("%s\n", data)
}
func SearchIssues(terms []string) (*IssuesSearchResult , error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" +q)
	if err != nil{
		return nil, err
	}
	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed : %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil{
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}