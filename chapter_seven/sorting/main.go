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
	Length int
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, 3},
	{"Go", "Moby", "Moby", 1992, 3},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, 4},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, 4},
}

func length(s string) int {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return int(d / time.Second)
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // вычисление размеров столбцов и вывод таблицы
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}
func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	fmt.Println("byArtist:")
	printTracks(tracks)
	fmt.Println()

	sort.Sort(sort.Reverse(byArtist(tracks)))
	fmt.Println("byArtist sorted:")
	printTracks(tracks)
	fmt.Println()
	palindrome := []int{1, 2, 3, 2, 1}
	fmt.Println("Is palindrome:", isPalindrome(byArtist(tracks)))
	fmt.Println("Is palindrome:", isPalindrome(sort.IntSlice(palindrome)))
}


func isPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}