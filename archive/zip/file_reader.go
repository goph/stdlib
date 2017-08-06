package zip

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

var (
	// ErrFileNotFound indicates that we reached the end of the file without finding the file
	ErrFileNotFound = errors.New("File not found in the archive")

	// ErrNotAFile indicates that there is a match for the path, but it's not a file (eg. it's a directory)
	ErrNotAFile = errors.New("Not a file")
)

// zipFileReader reads a certain file from a TAR archive
// which can also be optionally decompressed during the process
type zipFileReader struct {
	archive *zip.Reader

	found io.ReadCloser // The file object if it has been found
	file  string        // The path to the file inside the archive

	tmp *os.File
}

// NewZipFileReader returns a new Reader which reads a specific file from a .zip archive.
func NewZipFileReader(r io.Reader, f string) (io.ReadCloser, error) {
	tmp, err := ioutil.TempFile(os.TempDir(), "gitbin_")
	if err != nil {
		return nil, err
	}

	n, err := io.Copy(tmp, r)
	if err != nil {
		return nil, err
	}

	archive, err := zip.NewReader(tmp, n)
	if err != nil {
		return nil, err
	}

	return &zipFileReader{
		archive: archive,

		file: f,
		tmp:  tmp,
	}, nil
}

func (r *zipFileReader) Read(p []byte) (int, error) {
	// The file is already found, but hasn't been read entirely
	if r.found != nil {
		return r.found.Read(p)
	}

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.archive.File {
		if f.Name == r.file {
			fr, err := f.Open()
			if err != nil {
				return 0, err
			}

			r.found = fr

			return fr.Read(p)
		}
	}

	return 0, ErrFileNotFound
}

func (r *zipFileReader) Close() error {
	r.tmp.Close()
	os.Remove(r.tmp.Name())

	if r.found != nil {
		r.found.Close()
	}

	return nil
}
