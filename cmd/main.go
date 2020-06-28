package main

import (
	"fmt"
	"log"

	"github.com/adlandh/flickr-data-converter/src/flickr"
	"github.com/adlandh/flickr-data-converter/src/output"
	"github.com/adlandh/flickr-data-converter/src/settings"
)

func main() {

	mainSettings := settings.New()

	flickrData := flickr.New(mainSettings)

	if err := flickrData.Parse(); err != nil {
		log.Fatal(err)
	}

	outputObject := output.New(flickrData)

	if err := outputObject.Generatefolders(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All files successfully stored in " + mainSettings.Output)
}
