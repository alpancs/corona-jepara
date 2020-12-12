package main

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexHTML, _ := os.Open("index.html")
		defer indexHTML.Close()
		io.Copy(w, indexHTML)
	})

	http.HandleFunc("/chart_harian", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		URL := "https://corona.jepara.go.id/data/chart_harian"
		ctx := context.WithTimeout(5 * time.Second)
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
		req.Header.Set("Origin", "https://corona.jepara.go.id")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(getRawJSON(resp.Body))
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func getRawJSON(body io.ReadCloser) []byte {
	defer body.Close()
	bodyBytes, _ := ioutil.ReadAll(body)
	return bodyBytes[bytes.Index(bodyBytes, []byte("[{")):]
}
