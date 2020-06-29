package flickr

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adlandh/flickr-data-converter/src/settings"
)

var validFile = regexp.MustCompile(`photo_[0-9]+.json$`)

type Flickr struct {
	Settings       settings.Settings `json:"-"`
	Albums         []Album           `json:"albums"`
	Photos         map[string]Photo  `json:"-"`
	photoDataFiles []string
	photoFiles     []string
}

func New(settings settings.Settings) Flickr {
	return Flickr{
		Settings: settings,
		Photos:   make(map[string]Photo),
	}
}

func (f *Flickr) Parse() error {
	if err := f.getFiles(); err != nil {
		return err
	}
	if err := f.parseAlbums(); err != nil {
		return err
	}
	if err := f.parsePhotos(); err != nil {
		return err
	}
	return nil
}

func (f *Flickr) parseAlbums() error {
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

func (f *Flickr) parsePhotos() error {
	for _, file := range f.photoDataFiles {
		if !validFile.MatchString(file) {
			continue
		}

		photo, err := f.parsePhoto(file)

		if err != nil {
			return err
		}

		f.Photos[photo.Id] = photo
	}

	return nil
}

func (f *Flickr) getFiles() error {
	photoDataFiles, err := filepath.Glob(f.Settings.Data + string(os.PathSeparator) + PhotoFIles)
	if err != nil {
		return err
	}

	f.photoDataFiles = photoDataFiles

	photoFiles, err := filepath.Glob(f.Settings.Data + string(os.PathSeparator) + "*_o.jpg")
	if err != nil {
		return err
	}

	f.photoFiles = photoFiles

	return nil
}

func (f Flickr) parsePhoto(file string) (Photo, error) {

	var photo Photo

	photoFile, err := os.Open(file)

	if err != nil {
		return photo, err
	}

	defer photoFile.Close()

	if err = json.NewDecoder(photoFile).Decode(&photo); err != nil {
		return photo, err
	}

	if photo.FileName, err = f.findFilenameById(photo.Id); err != nil {
		return photo, err
	}
	return photo, nil
}

func (f *Flickr) findFilenameById(id string) (string, error) {

	for _, file := range f.photoFiles {
		if strings.Contains(file, id) {
			return file, nil
		}
	}

	return "", fmt.Errorf("photo with id %s not found", id)
}
