package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var httpClient = http.Client{Timeout: 10 * time.Second}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mainHTMLFile, _ := os.Open("main.html")
		defer mainHTMLFile.Close()
		io.Copy(w, mainHTMLFile)
	})

	http.HandleFunc("/chart_harian", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		resp, err := httpClient.Get("http://corona.jepara.go.id/data/chart_harian")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	})

	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":80"
	}
	fmt.Printf("listening \"%s\"\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
