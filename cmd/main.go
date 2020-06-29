package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adlandh/flickr-data-converter/src/models"
	"github.com/adlandh/flickr-data-converter/src/services/flickr"
	"github.com/adlandh/flickr-data-converter/src/services/output"
)

func main() {
	var dataFolder, outputFolder string

	flag.StringVar(&dataFolder, "data", ".", "Full path to folder with flickr data")
	flag.StringVar(&outputFolder, "output", "."+string(os.PathSeparator)+"output", "Full path to	output folder")
	flag.Parse()

	mainSettings := &models.Settings{
		Data:   dataFolder,
		Output: outputFolder,
	}

	flickrData := flickr.New(mainSettings)

	if err := flickrData.Parse(); err != nil {
		log.Fatal(err)
	}

	outputObject := output.New(mainSettings, flickrData.Albums, flickrData.Photos)

	if err := outputObject.Generatefolders(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All files successfully stored in " + mainSettings.Output)
}
