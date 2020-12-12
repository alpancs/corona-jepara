package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		resp, err := http.Get("https://corona.jepara.go.id/data/chart_harian")
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
