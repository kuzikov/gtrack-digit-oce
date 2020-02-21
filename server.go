package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"text/template"

	u "./utils"
)

// uninit template struct
var tpl *template.Template

// init initialize stuct tpl
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

//Creating global structures
// var (
// 	Tracker []u.Artist
// 	Dates   []u.ConcertDates
// 	Locs    []u.Locations
// 	Rels    []u.Relations
// )
// var Ticker = time.NewTicker(10 * time.Minute)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	go u.ApiFetch()

	// update info every 10 min.
	go u.ApiUpdate()

	http.HandleFunc("/", Index)

	http.HandleFunc("/search", Search)

	static := http.FileServer(http.Dir("public"))

	http.Handle("/public/", http.StripPrefix("/public/", static))

	http.ListenAndServe(":"+u.Port, nil)
}

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.html", &u.Tracker)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

// Search handler
func Search(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		text, err := url.ParseQuery(string(body))
		if err != nil {
			return
		}
		if !CheckKeys(text, "text") {
			errMess := u.ErrMess{Code: http.StatusBadRequest, Message: "can't reading keys from request"}
			ShowErr(w, &errMess)
			return
		}

		for i := range u.Tracker {
			if u.Tracker[i].Name == text.Get("text") {
				tempStruct := &u.Locations{}
				file, err := os.Open(u.Tracker[i].LocationsURL)
				if err != nil {
					fmt.Println(err.Error())
				}
				b, err := ioutil.ReadAll(file)
				if err != nil {
					fmt.Println(err.Error())
				}
				json.Unmarshal(b, &tempStruct)
				// file, _ = os.Open(Tracker[i].ConcertDatesURL)
				// b, _ = ioutil.ReadAll(file)
				// json.Unmarshal(b, &Tracker[i].ConcertDates)
				// file, _ = os.Open(Tracker[i].RelationsURL)
				// b, _ = ioutil.ReadAll(file)
				// json.Unmarshal(b, &Tracker[i].Relations)
				fmt.Println(tempStruct)
				// fmt.Println(Tracker[i].ConcertDates)
				// fmt.Println(Tracker[i].Relations)
				err = tpl.ExecuteTemplate(w, "card.html", &u.Tracker[i])
				if err != nil {
					fmt.Println(err.Error())
					return
				}
			}
		}
		errMess := u.ErrMess{Code: http.StatusOK, Message: "К сожалению информации по такому запросу не найдено"}
		ShowErr(w, &errMess)
		return
	default:
		errMess := u.ErrMess{Code: http.StatusBadRequest, Message: "can't read information from needed Method"}
		ShowErr(w, &errMess)
		return
	}
}

// CheckKeys handler
func CheckKeys(str url.Values, key ...string) bool {
	arr := []bool{}
	for _, v := range key {
		var a bool = false
		for i := range str {
			if v == i {
				if str[i][0] == "" {
					return false
				}
				a = !a

			}
		}
		arr = append(arr, a)
	}
	for _, v := range arr {
		if v == false {
			return false
		}
	}
	return true
}

// ShowErr handler
func ShowErr(w http.ResponseWriter, errMess *u.ErrMess) {

	w.WriteHeader(errMess.Code)

	err := tpl.ExecuteTemplate(w, "error.html", errMess)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func fFind(file, dir string) bool {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.Name() != file {
			continue
		} else {
			return true
		}
	}
	return false
}
