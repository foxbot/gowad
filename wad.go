package gowad

// Wad is a container for a map of files and their blob contents
type Wad struct {
	// Type indicates what kind of archive this package is. This
	// field is only relevant to actual game files.
	Type string
	// Files contains a mapping of file name -> blob contents
	Files map[string]([]byte)
}
