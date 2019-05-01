package location

import (
	"fmt"
	"sort"
)


// Gets location data of all riders/drivers within a
// given square


type Locations struct {
	Lat map[float64]int
	Long map[float64]int
	SortedLat []float64
	SortedLong []float64
}

type User int
const (
	DRIVER User = iota
	RIDER
)

func (loc *Locations) AllLocations(lat1 float64, long1 float64, lat2 float64, long2 float64) []int {
	fmt.Println("Welcome to Uber Location Service!")

	startIndex := 0
	endIndex := len(loc.SortedLat) - 1
	for i, v := range loc.SortedLat {
		if startIndex == 0 && v > lat1 {
			startIndex = i
		}

		if endIndex == len(loc.SortedLat) - 1 && v > lat2 {
			endIndex = i
		}
	}
	idsByLat := loc.SortedLat[startIndex : endIndex + 1]
	fmt.Println("ids by lat", idsByLat)
	startIndex = 0
	endIndex = len(loc.SortedLong) - 1
	for i, v := range loc.SortedLong {
		if startIndex == 0 && v > long1 {
			startIndex = i
		}

		if endIndex == len(loc.SortedLat) - 1 && v > long2 {
			endIndex = i
		}
	}
	idsByLong := make([]int, 0)
	for _, id := range loc.SortedLong[startIndex : endIndex + 1] {
		idsByLong = append(idsByLong, loc.Long[id])
	}
	fmt.Println("ids by long", idsByLong)

	var res []int
	for latId := range idsByLat {
		for longId := range idsByLong {
			if latId == longId {
				res = append(res, latId)
			}
		}
	}
	return res
}

func (loc *Locations) Add(id int, lat float64, long float64) {
	loc.Lat[lat] = id
	loc.SortedLat = append(loc.SortedLat, lat)
	loc.Long[long] = id
	loc.SortedLong = append(loc.SortedLong, long)

	sort.Float64s(loc.SortedLong)
	sort.Float64s(loc.SortedLat)
}

