package tar_test

import (
	"testing"

	stdtar "archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/goph/stdlib/archive/tar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTarGz(t *testing.T, fileName string, contents []byte) io.Reader {
	header := &stdtar.Header{
		Name:     fileName,
		Mode:     0640,
		Uid:      1000,
		Gid:      1000,
		Size:     int64(len(contents)),
		ModTime:  time.Unix(1491731729, 0),
		Typeflag: stdtar.TypeReg,
		Uname:    "test",
		Gname:    "test",
	}

	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)
	tr := stdtar.NewWriter(gz)

	tr.WriteHeader(header)
	tr.Write(contents)

	require.NoError(t, tr.Close(), "failed closing archive: %v")
	require.NoError(t, gz.Close(), "failed finishing compression: %v")

	return buf
}

func TestTarGzFileReader(t *testing.T) {
	fileName := "test.txt"
	contents := []byte("test")
	tgz := createTarGz(t, fileName, contents)

	reader, err := tar.NewTarGzFileReader(tgz, fileName)

	require.NoError(t, err, "cannot create file reader: %v")

	defer reader.Close()

	received, err := ioutil.ReadAll(reader)

	require.NoError(t, err, "cannot read file: %v")
	assert.Equal(t, contents, received)
}

func TestTarGzFileReader_NotFound(t *testing.T) {
	fileName := "test.txt"
	contents := []byte("test")
	tgz := createTarGz(t, fileName, contents)

	reader, err := tar.NewTarGzFileReader(tgz, "not_test.txt")

	require.NoError(t, err, "cannot create file reader: %v")

	defer reader.Close()

	_, err = ioutil.ReadAll(reader)

	assert.Equal(t, err, tar.ErrFileNotFound)
}

func ExampleNewTarGzFileReader() {
	tgz, _ := os.Open("testdata/test.tar.gz")

	reader, _ := tar.NewTarGzFileReader(tgz, "test.txt")
	contents, _ := ioutil.ReadAll(reader)
	fmt.Println(string(contents))
	// Output: test
}
