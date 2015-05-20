package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NotFound(id string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found", Id: id}); err != nil {
		log.Panic(err)
	}

}

func (myjournal *Myjournal) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	var entry string
	var journals Journals

	entry = vars["entryId"]
	journals, err = myjournal.Parse(entry)
	if (err != nil) && (err.Error() == "NotFound") {
		NotFound(entry, w)
		return
	}

	b, err := json.MarshalIndent(journals, "", "    ")
	if err != nil {
		log.Panicf("ERROR: encoding JSON: %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (mj *Myjournal) Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	var entry string
	var journals Journals
	var current Journals

	journals, err = myjournal.Parse("")

	entry = vars["entryId"]
	if entry == "" {
		entry = journals[0].Id
	}

	current, err = myjournal.Parse(entry)
	if (err != nil) && (err.Error() == "NotFound") {
		NotFound(entry, w)
		return
	}

	page := Page{Title: "Vitaliy's Food Journal", Navbar: journals, Content: current}
	renderTemplate(w, "dop", &page)
}
