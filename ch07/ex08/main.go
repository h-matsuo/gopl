package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

// type Track struct {
// 	Title  string
// 	Artist string
// 	Album  string
// 	Year   int
// 	Length time.Duration
// }
type Track map[int]string

var tracks = []*Track{
	{0: "Go", 1: "Delilah", 2: "From the Roots Up", 3: "2012", 4: "3m38s"},
	{0: "Go", 1: "Moby", 2: "Moby", 3: "1992", 4: "3m37s"},
	{0: "Go Ahead", 1: "Alicia Keys", 2: "As I Am", 3: "2007", 4: "4m36s"},
	{0: "Ready 2 Go", 1: "Martin Solveig", 2: "Smash", 3: "2011", 4: "4m24s"},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, (*t)[0], (*t)[1], (*t)[2], (*t)[3], (*t)[4])
	}
	tw.Flush()
}

func main() {
	sortKeys := []int{0, 3, 4}
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		for _, key := range sortKeys {
			if (*x)[key] != (*y)[key] {
				return (*x)[key] < (*y)[key]
			}
		}
		return false
	}})
	printTracks(tracks)
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
