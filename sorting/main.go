package sorting

import (
	"time"
	"text/tabwriter"
	"os"
	"fmt"
	"sort"
)

type Track struct {
	Title string
	Artis string
	Album string
	Year  int
	Lengtg time.Duration
}

type byArtis []*Track

func (x byArtis)Len()int  {
	return len(x)
}
func (x byArtis)Less(i,j int)bool  {
	return x[i].Artis<x[j].Artis
}
func (x byArtis)Swap(i,j int)  {
	x[i], x[j] = x[j], x[i]
}

var tracks  = []*Track{
	{"Go", "yangshen","nihao", 2012, Length("3m38s")},
	{"Go1", "yangshen1","nihao1", 2013, Length("3m38s")},
	{"Go2", "yangshen2","nihao2", 2014, Length("3m38s")},
}

func Length(s string) time.Duration  {
	d, err := time.ParseDuration(s)
	if err != nil{
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track)  {
	const format  = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ',0)
	fmt.Fprintf(tw, format, "Title", "Artis", "Album", "Year", "Lengtg")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "----", "------")
	sort.Sort(byArtis(tracks))
	for _,t :=range tracks{
		fmt.Fprintf(tw,format,t.Title,t.Artis,t.Album,t.Year,t.Lengtg)
	}
	tw.Flush()

}
