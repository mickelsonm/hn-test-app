package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mickelsonm/go-helpers/hackernews"
)

var (
	listenAddr = flag.Int("port", 3000, "http listen port")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/hackernews", func(w http.ResponseWriter, req *http.Request) {
		feed, err := hackernews.GetRssFeed()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		js, err := json.MarshalIndent(feed, " ", " ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Dude, it's an API.")
	})

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(fmt.Sprintf(":%d", *listenAddr))
}
