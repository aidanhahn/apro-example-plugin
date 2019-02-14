package main

import (
	"log"
	"net/http"
	"strconv"
)

func PluginMain(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["db"]

	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusOK)
		return
	}

	key := keys[0]
	keyInt, err := strconv.Atoi(key)

	if err != nil {
		log.Printf("param is not an int \n")
		w.WriteHeader(http.StatusOK)
		return
	}

	if keyInt%2 == 0 {
		w.Header().Set("Host", "even.domain2.nkrause.k736.net")
	} else {
		w.Header().Set("Host", "odd.domain2.nkrause.k736.net")
	}

	w.WriteHeader(http.StatusOK)
}
