package location

import (
	"fmt"
	"testing"
)

func TestLocations_Add(t *testing.T) {
	loc := Locations{Lat: make(map[float64]int), Long: make(map[float64]int)}
	loc.Add(12, 23.232, 89.232)
	if len(loc.SortedLong) != 1 {
		t.Error("Invalid operation")
	}
	fmt.Println(loc)
}

func TestLocations_AllLocations(t *testing.T) {
	loc := Locations{Lat: make(map[float64]int), Long: make(map[float64]int)}
	loc.Add(12, 23, 89)
	loc.Add(23, 24, 90)
	loc.Add(30, 28, 95)

	ids := loc.AllLocations(20, 80, 30, 100)
	fmt.Println(ids)
}
