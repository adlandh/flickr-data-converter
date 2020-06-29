package models

type Settings struct {
	Data   string
	Output string
}

func New(dataFolder, outputFolder string) Settings {
	return Settings{
		Data:   dataFolder,
		Output: outputFolder,
	}
}
