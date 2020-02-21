package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	TrackerUpd []Artist
	DatesUpd   []ConcertDates
	LocsUpd    []Locations
	RelsUpd    []Relations
)

func ApiUpdate() {
	for {
		time.Sleep(10 * time.Minute)
		resp1, _ := http.Get(URLartists)
		array1, _ := ioutil.ReadAll(resp1.Body)
		resp2, _ := http.Get(URLdates)
		array2, _ := ioutil.ReadAll(resp2.Body)
		resp3, _ := http.Get(URLlocations)
		array3, _ := ioutil.ReadAll(resp3.Body)
		resp4, _ := http.Get(URLrelation)
		array4, _ := ioutil.ReadAll(resp4.Body)
		array2 = array2[9 : len(array2)-2]
		array3 = array3[9 : len(array3)-2]
		array4 = array4[9 : len(array4)-2]
		json.Unmarshal(array1, &TrackerUpd)
		json.Unmarshal(array2, &DatesUpd)
		json.Unmarshal(array3, &LocsUpd)
		json.Unmarshal(array4, &RelsUpd)

		// fmt.Printf("lens of arrays: %v , %v, %v, %v \n", len(Tracker), len(Dates), len(Locs), len(Rels))

		fmt.Println(DatesUpd[1])
		fmt.Println(LocsUpd[1])
		fmt.Println(RelsUpd[1])
		for i := range Tracker {
			TrackerUpd[i].ConcertDates = DatesUpd[i]
			TrackerUpd[i].Locations = LocsUpd[i]
			TrackerUpd[i].Locations.Dates = DatesUpd[i]
			TrackerUpd[i].Relations = RelsUpd[i]
		}
		Tracker = TrackerUpd
	}
}

func ApiFetch() {
	resp1, _ := http.Get(URLartists)
	array1, _ := ioutil.ReadAll(resp1.Body)
	resp2, _ := http.Get(URLdates)
	array2, _ := ioutil.ReadAll(resp2.Body)
	resp3, _ := http.Get(URLlocations)
	array3, _ := ioutil.ReadAll(resp3.Body)
	resp4, _ := http.Get(URLrelation)
	array4, _ := ioutil.ReadAll(resp4.Body)
	array2 = array2[9 : len(array2)-2]
	array3 = array3[9 : len(array3)-2]
	array4 = array4[9 : len(array4)-2]
	json.Unmarshal(array1, &Tracker)
	json.Unmarshal(array2, &Dates)
	json.Unmarshal(array3, &Locs)
	json.Unmarshal(array4, &Rels)

	// fmt.Printf("lens of arrays: %v , %v, %v, %v \n", len(Tracker), len(Dates), len(Locs), len(Rels))

	fmt.Println(Dates[1])
	fmt.Println(Locs[1])
	fmt.Println(Rels[1])
	for i := range Tracker {
		Tracker[i].ConcertDates = Dates[i]
		Tracker[i].Locations = Locs[i]
		Tracker[i].Locations.Dates = Dates[i]
		Tracker[i].Relations = Rels[i]
	}

}
