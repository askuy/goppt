package main

import (
	"flag"
	"log"

	"golang.org/x/tools/present"

	"net/http"

	"github.com/askuy/goppt/libs"
)

var (
	httpAddr   = flag.String("http", "127.0.0.1:3999", "HTTP service address (e.g., '127.0.0.1:3999')")
	originHost = flag.String("orighost", "", "host component of web origin URL (e.g., 'localhost')")
	basePath   = flag.String("base", "", "base path for slide template and static resources")
)

func main() {
	flag.BoolVar(&present.PlayEnabled, "play", true, "enable playground (permit execution of arbitrary user code)")
	flag.BoolVar(&present.NotesEnabled, "notes", false, "enable presenter notes (press 'N' from the browser to display them)")
	flag.Parse()
	log.Print("strings")
	if *basePath == "" {
	}

	err := libs.InitTemplates(*basePath)
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
	libs.Init()

	http.ListenAndServe(*httpAddr, nil) //outprint hello world！

}
