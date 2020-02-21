package utils

var (
	Tracker []Artist
	Dates   []ConcertDates
	Locs    []Locations
	Rels    []Relations
)

type Artist struct {
	Id              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsURL    string   `json:"locations"`
	ConcertDatesURL string   `json:"concertDates"`
	RelationsURL    string   `json:"relations"`
	Locations       Locations
	ConcertDates    ConcertDates
	Relations       Relations
}

type Locations struct {
	Id        int          `json:"id"`
	Locations []string     `json:"locations"`
	Dates     ConcertDates `json:"concertDates"`
}

type ConcertDates struct {
	// Index []struct {
	// 	Id    int      `json:"id"`
	// 	Dates []string `json:"dates"`
	// } `json:"index"`
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

//123
const (
	URLartists   = "https://groupietrackers.herokuapp.com/api/artists"
	URLlocations = "https://groupietrackers.herokuapp.com/api/locations"
	URLdates     = "https://groupietrackers.herokuapp.com/api/dates"
	URLrelation  = "https://groupietrackers.herokuapp.com/api/relation"
	// URLartists   = "./public/api/artists.json"
	// URLlocations = "./public/api/locations.json"
	// URLdates     = "./public/api/dates.json"
	// URLrelation  = "./public/api/relation.json"
)

const (
	Port = "80"
)

type ErrMess struct {
	Code    int
	Message string
}
