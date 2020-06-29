package output

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/adlandh/flickr-data-converter/src/models"
)

type Output struct {
	Settings *models.Settings
	Albums   models.Albums
	Photos   models.Photos
	albumsWg *sync.WaitGroup
}

func (o *Output) Generatefolders() {
	for _, album := range o.Albums {
		o.albumsWg.Add(1)
		go o.generateAlbumFolder(album)
	}

	o.albumsWg.Wait()
}

func New(settings *models.Settings, albums models.Albums, photos models.Photos) Output {
	return Output{
		Settings: settings,
		Albums:   albums,
		Photos:   photos,
		albumsWg: &sync.WaitGroup{},
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

func (o *Output) generateAlbumFolder(album models.Album) {
	defer o.albumsWg.Done()

	albumDir := o.Settings.Output + string(os.PathSeparator) + album.Title
	err := os.Mkdir(albumDir, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(models.NewOutputError(albumDir, "", err))
	}

	for _, photoID := range album.Photos {
		if photoID == "0" {
			continue
		}

		srcFile := o.Photos[photoID].FileName
		dstFile := albumDir + string(os.PathSeparator) + o.Photos[photoID].Name + ".jpg"
		_, err := o.copy(srcFile, dstFile)
		if err != nil {
			log.Fatal(models.NewOutputError(albumDir, photoID, err))
		}
	}
}
