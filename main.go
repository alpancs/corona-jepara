package main

import (
	"fmt"
	"io"
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
		resp, err := httpClient.Get("https://corona.jepara.go.id/data/chart_harian")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	})

	addr := ":" + os.Getenv("PORT")
	fmt.Printf("listening \"%s\"\n", addr)
	http.ListenAndServe(addr, nil)
}
