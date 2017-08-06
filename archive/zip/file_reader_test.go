package zip_test

import (
	"testing"

	stdzip "archive/zip"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/goph/stdlib/archive/zip"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createZip(t *testing.T, fileName string, contents []byte) *os.File {
	tmp, err := ioutil.TempFile(os.TempDir(), "gitbin_")
	require.NoError(t, err, "error creating temp file")

	w := stdzip.NewWriter(tmp)

	f, err := w.Create(fileName)
	require.NoError(t, err)

	_, err = f.Write(contents)
	require.NoError(t, err)

	require.NoError(t, w.Close(), "failed closing archive")

	tmp.Close()

	tmp, err = os.Open(tmp.Name())
	require.NoError(t, err)

	return tmp
}

func cleanupZip(f *os.File) {
	f.Close()
	os.Remove(f.Name())
}

func TestZipFileReader(t *testing.T) {
	fileName := "test.txt"
	contents := []byte("test")
	z := createZip(t, fileName, contents)

	reader, err := zip.NewZipFileReader(z, fileName)

	require.NoError(t, err, "cannot create file reader")

	defer reader.Close()

	received, err := ioutil.ReadAll(reader)

	require.NoError(t, err, "cannot read file")
	assert.Equal(t, contents, received)
}

func TestZipFileReader_NotFound(t *testing.T) {
	fileName := "test.txt"
	contents := []byte("test")
	z := createZip(t, fileName, contents)

	reader, err := zip.NewZipFileReader(z, "not_test.txt")

	require.NoError(t, err, "cannot create file reader")

	defer reader.Close()

	_, err = ioutil.ReadAll(reader)

	assert.Equal(t, err, zip.ErrFileNotFound)
}

func ExampleNewZipFileReader() {
	f, _ := os.Open("testdata/test.zip")

	reader, _ := zip.NewZipFileReader(f, "test.txt")
	contents, _ := ioutil.ReadAll(reader)
	fmt.Println(string(contents))
	// Output: test
}
