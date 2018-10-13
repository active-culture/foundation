package main

// the import-ant styff
import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "time"
)

// the schema
type Coob struct {
    ID        string `json:"id"`
    Type      string `json:"type,omitempty"`
    Name      string `json:"name,omitempty"`
    Avatar    string `json:"avatar,omitempty"`
}

type Event struct {
    ID        string `json:"id,omitempty"`
    Type      string `json:"type,omitempty"`
    Name      string `json:"name,omitempty"`
    Image     string `json:"image,omitempty"`
    Info      string `json:"blurb,omitempty"`
    Location  string `json:"location,omitempty"`
    Link      string `json:"link,omitempty"`
    Time      time.Time
}

// the databae
var coobs []Coob
var events []Event

// api
func GetCoob(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(coobs)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(events)
}

// the main squeeze
func main() {
    coobs = append(coobs, Coob{Name: "Coobmaster General", ID: "0", Type:"Squalid", Avatar:"http://activeculture.nyc/img/Keroro_1.png"})
    coobs = append(coobs, Coob{Name: "Coob Garou", ID: "1", Type:"Swamp-esque", Avatar:"http://activeculture.nyc/img/Terminator-1984_612x380_0.jpg"})
    events = append(events, Event{Name: "Big Money Millions", ID: "0", Info: "Yea thats that ish"})
    router := mux.NewRouter()
    router.HandleFunc("/coob", GetCoob).Methods("GET")
    router.HandleFunc("/events", GetEvents).Methods("GET")
    fmt.Println("Done")
    fmt.Println(coobs)
    log.Fatal(http.ListenAndServe(":8081", router))
}

// 🙏 to @codehakase for the article
