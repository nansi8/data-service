package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/data-service/health", health)
	r.HandleFunc("/data-service/data-nodes", DataNodesHandler)
	r.HandleFunc("/data-service/checksum-nodes", ChecksumNodesHandler)

	http.ListenAndServe(":80", r)
}

func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func DataNodesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(Config().dataNodes)))
}

func ChecksumNodesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(Config().checksumNodes)))
}
