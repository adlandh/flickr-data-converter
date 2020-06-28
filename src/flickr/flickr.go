package flickr

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adlandh/flickr-data-converter/src/settings"
)

var validFile = regexp.MustCompile(`photo_[0-9]+.json$`)

type Flickr struct {
	Settings settings.Settings `json:"-"`
	Albums   []Album           `json:"albums"`
	Photos   map[string]Photo  `json:"-"`
}

func New(settings settings.Settings) Flickr {
	return Flickr{
		Settings: settings,
		Photos:   make(map[string]Photo),
	}
}

func (f *Flickr) Parse() error {
	if err := f.ParseAlbums(); err != nil {
		return err
	}
	if err := f.ParsePhotos(); err != nil {
		return err
	}
	return nil
}

func (f *Flickr) ParseAlbums() error {
	albumFile, err := os.Open(f.Settings.Data + string(os.PathSeparator) + AlbumFileName)
	if err != nil {
		return err
	}

	defer albumFile.Close()

	if err = json.NewDecoder(albumFile).Decode(f); err != nil {
		return err
	}
	return nil
}

func (f *Flickr) ParsePhotos() error {

	photoFiles, err := filepath.Glob(f.Settings.Data + string(os.PathSeparator) + PhotoFIles)
	if err != nil {
		return err
	}

	for _, file := range photoFiles {
		if !validFile.MatchString(file) {
			continue
		}

		photo, err := f.ParsePhoto(file)

		if err != nil {
			return err
		}

		f.Photos[photo.Id] = photo
	}

	return nil

}

func (f Flickr) ParsePhoto(file string) (Photo, error) {

	var photo Photo

	photoFile, err := os.Open(file)

	if err != nil {
		return photo, err
	}

	defer photoFile.Close()

	if err = json.NewDecoder(photoFile).Decode(&photo); err != nil {
		return photo, err
	}

	name := strings.ReplaceAll(photo.Name, " ", "-")
	name = strings.ReplaceAll(name, ".", "")
	name = strings.ReplaceAll(name, "(", "")
	name = strings.ReplaceAll(name, ")", "")

	photo.FileName = strings.ToLower(name) + "_" + photo.Id + "_o.jpg"

	return photo, nil
}
