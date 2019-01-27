package gowad

import (
	"errors"
)

// ErrInvalidHeader occurs when an archive is passed which contains an invalid
// or corrupt header, which is one less than 12 bytes in length
var ErrInvalidHeader = errors.New("archive header is too short, must be invalid")

// ErrInvalidDirectory occurs when an archive contains a directory which is not
// a multiple of 16 bytes in length
var ErrInvalidDirectory = errors.New("archive directory is improperly sized")
