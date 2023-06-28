package main

import (
	"encoding/json"
	"net/http"
)

type Course struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Id string `json:"id"`
	Author *Author `json:"author"`
} 

type Author struct {
	Name string `json:"name"`
	Github string `json:"github"`
}

// A fake database
var courseDB []Course

func (c *Course) IsEmpty() bool {
	return c.Name == "" && c.Id == ""
}

func main()  {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	if len(courseDB) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(courseDB)
}

func addCourse(w http.ResponseWriter, r *http.Request) {}

func getCourse(w http.ResponseWriter, r *http.Request) {}

func updateCourse(w http.ResponseWriter, r *http.Request) {}

func deleteCourse(w http.ResponseWriter, r *http.Request) {}



func courses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCourses(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func course(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
	case "PUT":
	case "DELETE":
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}





