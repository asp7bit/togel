package main

import (
	"log"
	"fmt"
	"net/http"

	"github.com/drone/routes"
	"github.com/codegangsta/negroni"
	"github.com/rtulus/togel/src/conf"
	"github.com/rtulus/togel/src/nomor"
)

var config conf.Config

func readConfig() {
	
	fileLocation := fmt.Sprintf("files/togel/devel.ini")
	log.Println(fmt.Sprintf("Using configuration : %s", fileLocation))
	ok := conf.ReadConfig(&config, fileLocation)

	if !ok {
		log.Fatal("Could not find configuration file")
	}
	log.Println(config.TkpCore.SlaveDB)
	conf.InitDB(&config.TkpCore.SlaveDB)
}

func main() {
	log.Println("Hello World");
	
	readConfig()

	mux := routes.New()
	mux.Post("/togel/new", nomor.New)
	
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":3434", n)

	
}
