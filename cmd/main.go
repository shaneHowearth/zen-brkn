// Package main -
package main

import (
	"log"

	"github.com/shanehowearth/zen/repository/jsonfiles"
	simpleSearch "github.com/shanehowearth/zen/search/simple"
	simpleUI "github.com/shanehowearth/zen/ui/simple"
	zensearch "github.com/shanehowearth/zen/zen"
)

func main() {
	// Create the concrete instances required to run the app

	// Repository
	ds, err := jsonfiles.NewDatastore(map[string]string{
		"Organizations": "../repository/jsonfiles/data/organizations.json",
		"Tickets":       "../repository/jsonfiles/data/tickets.json",
		"Users":         "../repository/jsonfiles/data/users.json",
	})

	if err != nil {
		log.Fatalf("Datastore creation error %v", err)
	}

	// UI
	ui := simpleUI.Simple{}

	// Search
	s := simpleSearch.Simple{}

	// Tie them all together and start the application
	z, err := zensearch.NewZen(ds, ui, s)
	if err != nil {
		log.Fatalf("Unable to start with error %v", err)
	}
	z.Run()
}
