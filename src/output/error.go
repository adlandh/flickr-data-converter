package output

import (
	"fmt"
)

type OutputError struct {
	Album string
	Photo string
	Err   error
}

func NewError(album, photo string, err error) OutputError {
	return OutputError{
		Album: album,
		Photo: photo,
		Err:   err,
	}
}

func (o OutputError) Error() string {
	return fmt.Sprintf("Album: %s, photo: %s, err: %s", o.Album, o.Photo, o.Err.Error())
}
