package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/data-service", index)
	r.HandleFunc("/data-service/data-nodes", DataNodesHandler)
	r.HandleFunc("/data-service/checksum-nodes", ChecksumNodesHandler)

	http.ListenAndServe(":80", r)
}

func DataNodesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(Config().dataNodes)))
}

func ChecksumNodesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(Config().checksumNodes)))
}

func index(w http.ResponseWriter, req *http.Request) {
	// all calls to unknown url paths should return 404
	if req.URL.Path != "/" {
		log.Printf("404: %s", req.URL.String())
		http.NotFound(w, req)
	} else {
		w.Write([]byte("Test\n"))
	}
}
