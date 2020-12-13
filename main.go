package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var httpClient = http.Client{Timeout: 5 * time.Second}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mainHTML, _ := ioutil.ReadFile("main.html")
		w.Write(mainHTML)
	})

	http.HandleFunc("/chart_harian", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.Redirect(w, r, "https://corona.jepara.go.id/data/chart_harian", http.StatusFound)
	})

	addr := ":" + os.Getenv("PORT")
	fmt.Printf("listening \"%s\"\n", addr)
	http.ListenAndServe(addr, nil)
}
