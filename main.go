package main

import (
	"bytes"
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
		req, _ := http.NewRequest(http.MethodGet, "https://cors-anywhere.herokuapp.com/https://corona.jepara.go.id/data/chart_harian", nil)
		req.Header.Set("Origin", "https://corona.jepara.go.id")
		resp, err := httpClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(getRawJSON(resp.Body))
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func getRawJSON(body io.ReadCloser) []byte {
	defer body.Close()
	bodyBytes, _ := ioutil.ReadAll(body)
	return bodyBytes[bytes.Index(bodyBytes, []byte("[{")):]
}
