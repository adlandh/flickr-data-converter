package output

import (
	"fmt"
	"io"
	"os"

	"github.com/adlandh/flickr-data-converter/src/flickr"
)

type Output struct {
	flickr.Flickr
}

func (o Output) Generatefolders() error {
	for _, album := range o.Albums {
		albumDir := o.Settings.Output + string(os.PathSeparator) + album.Title
		err := os.Mkdir(albumDir, 0755)
		if err != nil && !os.IsExist(err) {
			return NewError(albumDir, "", err)
		}
		for _, photoId := range album.Photos {
			if photoId == "0" {
				continue
			}
			srcFile := o.Settings.Data + string(os.PathSeparator) + o.Photos[photoId].FileName
			dstFile := albumDir + string(os.PathSeparator) + o.Photos[photoId].Name + ".jpg"
			_, err := Copy(srcFile, dstFile)
			if err != nil {
				return NewError(albumDir, photoId, err)
			}
		}
	}

	return nil
}

func New(flickr2 flickr.Flickr) Output {
	return Output{
		flickr2,
	}
}

func Copy(src, dst string) (int64, error) {
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