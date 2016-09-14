package main

import (
	"log"
#	"github.com/goji/goji"
#	"github.com/rtulus/togel/src/nomor"
)

func main() {
	log.Println("Hello World");

#	mux := goji.NewMux()
#	mux.HandleFuncC(pat.Get("togel/new"), nomor.New)

#	http.ListenAndServe("localhost:3434", mux)
}


