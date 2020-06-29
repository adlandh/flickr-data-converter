package output

import (
	"fmt"
	"io"
	"os"

	"github.com/adlandh/flickr-data-converter/src/models"
)

type Output struct {
	settings *models.Settings
	Albums   models.Albums
	Photos   models.Photos
}

func (o Output) Generatefolders() error {
	for _, album := range o.Albums {
		albumDir := o.settings.Output + string(os.PathSeparator) + album.Title
		err := os.Mkdir(albumDir, 0755)
		if err != nil && !os.IsExist(err) {
			return models.NewOutputError(albumDir, "", err)
		}
		for _, photoId := range album.Photos {
			if photoId == "0" {
				continue
			}
			srcFile := o.Photos[photoId].FileName
			dstFile := albumDir + string(os.PathSeparator) + o.Photos[photoId].Name + ".jpg"
			_, err := o.copy(srcFile, dstFile)
			if err != nil {
				return models.NewOutputError(albumDir, photoId, err)
			}
		}
	}

	return nil
}

func New(settings *models.Settings, albums models.Albums, photos models.Photos) Output {
	return Output{
		settings: settings,
		Albums:   albums,
		Photos:   photos,
	}
}

func (o Output) copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
