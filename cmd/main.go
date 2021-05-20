// Package main -
package main

import (
	"log"

	"github.com/shanehowearth/zen/repo"
	"github.com/shanehowearth/zen/ui"
	"github.com/shanehowearth/zen/zen"
)

func main() {
	z, err := zen.NewBrain(&repo.Data{}, &ui.CLI{})
	if err != nil {
		log.Fatalf("Dying, because creating a new brain died with error: %v", err)
	}
	z.Forever()
}
