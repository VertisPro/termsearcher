package main

import (
	"bitbucket.org/harshadp/ontotestpad/src/common"
	"bitbucket.org/harshadp/ontotestpad/src/handlers"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/phyber/negroni-gzip/gzip"
	"net/http"
)

func main() {
	common.Initvars()
	common.Setupdb()
	// common.Setupspellchecker()
	// TODO: Setup shutdown when server is given a signal - see common.SetupShutdown
	defer common.Shutdown()
	// fmt.Println(common.Spk.SuggestSpelling("artointestinal"))
	// fmt.Println(common.Spk.CheckText("Myocariel Infraction"))

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(common.Siteroot)))
	mux.HandleFunc("/snomed/searchterm", handlers.Snomedsearchterm)
	mux.HandleFunc("/loinc/searchterm", handlers.Loincsearchterm)
	n.UseHandler(mux)

	fmt.Println("Ontoserver Started")
	n.Run(":80")
}
