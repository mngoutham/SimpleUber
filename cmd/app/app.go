package main

import (
	"fmt"
	"github.com/mngoutham/SimpleUber/pkg/location"
)

func main() {
	loc := location.Locations{Lat: make(map[float64]int), Long: make(map[float64]int)}
	loc.Add(12, 23.232, 89.232)
	if len(loc.SortedLong) != 1 {
		fmt.Errorf("Invalid operation")
	}
	fmt.Println(loc)

	//a := struct {name []string} {name:[]string{"1","2"}}
}
