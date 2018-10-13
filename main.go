package main

// the import-ant styff
import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

// the schema
type Coob struct {
    ID   string `json:"id,omitempty"`
    Type string `json:"type"`
    Name string `json:"name,omitempty"`
}

// the databae
var coobs []Coob

// api
func GetCoob(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(coobs)
}

// the main squeeze
func main() {
    coobs = append(coobs, Coob{Name: "Coobmaster General", ID: "0", Type:"Squalid"})
    router := mux.NewRouter()
    router.HandleFunc("/coob", GetCoob).Methods("GET")
    fmt.Println("Done")
    fmt.Println(coobs)
    log.Fatal(http.ListenAndServe(":8081", router))
}

// 🙏 to codehakase for the article
