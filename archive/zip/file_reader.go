//+build experimental

package zip

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

var (
	// ErrFileNotFound indicates that we reached the end of the archive without finding the file.
	ErrFileNotFound = errors.New("file not found in the archive")
)

// zipFileReader reads a certain file from a ZIP archive.
type zipFileReader struct {
	archive *zip.Reader
	file    string // The path to the file inside the archive.

	found io.ReadCloser // The file object if it has been found.

	tmp *os.File // The temporary file the archive is saved to.
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
		file:    f,

		tmp: tmp,
	}, nil
}

func (r *zipFileReader) Read(p []byte) (int, error) {
	// The file is already found
	if r.found != nil {
		return r.found.Read(p)
	}

	// Iterate through the files in the archive
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
