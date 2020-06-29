package flickr

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adlandh/flickr-data-converter/src/models"
)

var validFile = regexp.MustCompile(`photo_[0-9]+.json$`)

type Flickr struct {
	settings       models.Settings
	Albums         models.Albums
	Photos         models.Photos
	photoDataFiles []string
	photoFiles     []string
}

func New(settings models.Settings) Flickr {
	return Flickr{
		settings: settings,
		Photos:   make(models.Photos),
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
	var albums models.AlbumList
	albumFile, err := os.Open(f.settings.Data + string(os.PathSeparator) + models.AlbumFileName)
	if err != nil {
		return err
	}

	defer albumFile.Close()

	if err = json.NewDecoder(albumFile).Decode(&albums); err != nil {
		return err
	}

	f.Albums = albums.Albums

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
	photoDataFiles, err := filepath.Glob(f.settings.Data + string(os.PathSeparator) + models.PhotoDataFiles)
	if err != nil {
		return err
	}

	f.photoDataFiles = photoDataFiles

	photoFiles, err := filepath.Glob(f.settings.Data + string(os.PathSeparator) + "*_o.jpg")
	if err != nil {
		return err
	}

	f.photoFiles = photoFiles

	return nil
}

func (f Flickr) parsePhoto(file string) (models.Photo, error) {

	var photo models.Photo

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
