package settings

import (
	"flag"
	"os"
)

type Settings struct {
	Data   string
	Output string
}

func New() Settings {
	var dataFolder, outputFolder string
	flag.StringVar(&dataFolder, "data", ".", "Full path to folder with flickr data")
	flag.StringVar(&outputFolder, "output", "."+string(os.PathSeparator)+"output", "Full path to	output folder")

	flag.Parse()

	return Settings{
		Data:   dataFolder,
		Output: outputFolder,
	}
}
