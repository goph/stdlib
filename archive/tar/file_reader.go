package tar

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
)

var (
	// ErrFileNotFound indicates that we reached the end of the archive without finding the file..
	ErrFileNotFound = errors.New("file not found in the archive")

	// ErrNotAFile indicates that there is a match for the path, but it's not a file (eg. it's a directory).
	ErrNotAFile = errors.New("not a file")
)

// tarFileReader reads a certain file from a TAR archive
// which can also optionally be decompressed during the process.
type tarFileReader struct {
	archive *tar.Reader
	file    string // The path to the file inside the archive.

	found bool // Whether the file has been found or not.

	decompressor io.Closer // When there is decompression involved, the decompressor might have to be closed.
}

// NewTarGzFileReader returns a new Reader which reads a specific file from a .tar.gz archive.
func NewTarGzFileReader(r io.Reader, f string) (io.ReadCloser, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &tarFileReader{
		decompressor: gz,
		archive:      tar.NewReader(gz),

		file: f,
	}, nil
}

func (r *tarFileReader) Read(p []byte) (int, error) {
	// The file is already found
	if r.found {
		return r.archive.Read(p)
	}

	// Iterate through the files in the archive
	for {
		header, err := r.archive.Next()
		if err == io.EOF {
			return 0, ErrFileNotFound
		} else if err != nil {
			return 0, err
		}

		if header.Name == r.file {
			switch header.Typeflag {
			case tar.TypeReg, tar.TypeRegA:
				r.found = true
				return r.archive.Read(p)
			default: // Avoid scanning the rest of the archive
				return 0, ErrNotAFile
			}
		}
	}
}

func (r *tarFileReader) Close() error {
	if r.decompressor != nil {
		return r.decompressor.Close()
	}

	return nil
}
