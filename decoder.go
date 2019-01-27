package gowad

import (
	"io"
)

const headerLen = 12

// ChunkSize states how many bytes should be read from the archive at a time.
var ChunkSize = 256

// Decode reads a WAD archive from an io.Reader
func Decode(r io.Reader) (*Wad, error) {
	header := make([]byte, headerLen)
	n, err := r.Read(header)
	if err != nil {
		return nil, err
	}
	if n < headerLen {
		return nil, ErrInvalidHeader
	}

	format := string(header[0:4])
	dirLen := readInt(header[4:8])
	dirOffset := readInt(header[8:12])
	print(format)
	print(dirLen)
	print(dirOffset)

	data := []byte{}
	chunk := make([]byte, ChunkSize)
	for n > 0 {
		n, err = r.Read(chunk)
		if err != nil {
			return nil, err
		}
		data = append(data, chunk[:n]...)
	}

	actualDirLen := len(data[dirOffset : dirOffset+dirLen])
	if actualDirLen%16 != 0 || actualDirLen != dirLen {
		return nil, ErrInvalidDirectory
	}
	dir := make(map[string]([]byte))

	for i := 0; i < dirLen/16; i++ {
		offset := dirOffset + (i * 16)
		entry := data[offset : offset+16]
		fileOffset := readInt(entry[0:4])
		fileSize := readInt(entry[4:8])
		fileName := readString(entry[8:16])
		dir[fileName] = data[fileOffset : fileOffset+fileSize]
	}

	w := Wad{
		Type:  format,
		Files: dir,
	}

	return &w, nil
}

func readString(buf []byte) string {
	// note: zDoom wiki claims that we should null-terminate, golang will just
	// ignore null characters. this might be an inconsistency if some archive
	// contains weird data after a null-terminated file name
	return string(buf)
}

func readInt(buf []byte) int {
	return (int(buf[0]) << 24) | (int(buf[1]) << 16) | (int(buf[2]) << 8) | int(buf[3])
}
