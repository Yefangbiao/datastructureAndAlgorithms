package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	cus := customSort{t: tracks}
	cus.addSortBy("Artist")
	sort.Sort(&cus)
	printTracks(cus.t)
	cus.addSortBy("Album")
	sort.Sort(&cus)
	printTracks(cus.t)
	cus.addSortBy("Year")
	sort.Sort(&cus)
	printTracks(cus.t)
}

type customSort struct {
	t     []*Track
	multi []string
}

func (x *customSort) Len() int { return len(x.t) }
func (x *customSort) Less(i, j int) bool {
	for _, str := range x.multi {
		switch str {
		case "Title":
			return x.t[i].Title < x.t[j].Title
		case "Artist":
			return x.t[i].Artist < x.t[j].Artist
		case "Album":
			return x.t[i].Album < x.t[j].Album
		case "Year":
			return x.t[i].Year < x.t[j].Year
		case "Length":
			return x.t[i].Length < x.t[j].Length
		default:
			return false
		}
	}
	return false
}
func (x *customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x *customSort) addSortBy(sortType string) {
	newMulti := make([]string, 0, len(x.multi)+1)
	newMulti = append(newMulti, sortType)
	newMulti = append(newMulti, x.multi...)
	x.multi = newMulti
}
