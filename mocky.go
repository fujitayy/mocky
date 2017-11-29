package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	if "/" == req.URL.Path {
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			fmt.Errorf("%s", err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}

		io.WriteString(w, string(content))
	}
}

func DataHandler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1]

	d := rand.Intn(10)
	time.Sleep(time.Duration(d) * time.Second)

	content, err := ioutil.ReadFile(filepath.Join("data", fileName))
	if err != nil {
		fmt.Errorf("%s", err.Error())
		http.NotFound(w, req)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, string(content))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/data/", DataHandler)

	signal := make(chan bool)

	go func() {
		log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
		close(signal)
	}()

	fmt.Println("Server listening on 127.0.0.1:8080")
	<-signal
}
