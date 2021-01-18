// Package main -
package main

import (
	"log"

	"github.com/shanehowearth/zen/repository/jsonfiles"
	simpleSearch "github.com/shanehowearth/zen/search/simple"
	"github.com/shanehowearth/zen/ui/barebones"
	zensearch "github.com/shanehowearth/zen/zen"
)

func main() {
	// Create the concrete instances required to run the app

	// Repository
	// Note: There is no check on number of files, or that certain files have
	// been added. Only a check that at least one has been added.
	// Having no checks allows any number of other sets to be added.
	ds, err := jsonfiles.NewDatastore(map[string]string{
		"Organizations": "../repository/jsonfiles/data/organizations.json",
		"Tickets":       "../repository/jsonfiles/data/tickets.json",
		"Users":         "../repository/jsonfiles/data/users.json",
	})

	if err != nil {
		log.Fatalf("Datastore creation error %v", err)
	}

	// UI
	ui := barebones.Bb{}

	// Search
	s := simpleSearch.Searcher{}

	// Tie them all together and start the application
	z, err := zensearch.NewZen(ds, ui, s)
	if err != nil {
		log.Fatalf("Unable to start with error %v", err)
	}
	z.Run()
}
