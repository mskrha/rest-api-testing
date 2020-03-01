package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	LINE = "================================================================================"
	PORT = "12345"
)

func debug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, time.Now().Format("2006-01-02 15:04:05.000000000")+": "+format, a...)
}

func line() {
	fmt.Println(LINE)
}

func main() {
	fmt.Println("REST API testing server")
	debug("PID: %d\n", os.Getpid())
	http.HandleFunc("/", handle)
	debug("Listening for HTTP queries on port %s ...\n", PORT)
	line()
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		panic(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	debug("Client IP: %s\n", r.RemoteAddr)
	debug("Method: %s\n", r.Method)
	debug("URL: %s\n", r.URL)
	for k, v := range r.Header {
		for _, x := range v {
			debug("Header: %s: %s\n", k, x)
		}
	}
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		b := string(body)
		if len(b) > 0 {
			debug("Body: %s\n", b)
		} else {
			debug("Empty body\n")
		}
	} else {
		debug("Failed to read body\n")
	}
	line()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	return
}
