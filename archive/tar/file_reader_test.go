package tar

import (
	"testing"

	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func createTarGz(t *testing.T, fileName string, contents []byte) io.Reader {
	header := &tar.Header{
		Name:     fileName,
		Mode:     0640,
		Uid:      1000,
		Gid:      1000,
		Size:     int64(len(contents)),
		ModTime:  time.Unix(1491731729, 0),
		Typeflag: tar.TypeReg,
		Uname:    "test",
		Gname:    "test",
	}

	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)
	tr := tar.NewWriter(gz)

	tr.WriteHeader(header)
	tr.Write(contents)

	err := tr.Close()
	if err != nil {
		t.Fatalf("failed closing archive: %v", err)
	}

	err = gz.Close()
	if err != nil {
		t.Fatalf("failed finishing compression: %v", err)
	}

	return buf
}

func TestTarGzFileReader(t *testing.T) {
	fileName := "test.txt"
	contents := []byte("test")
	tgz := createTarGz(t, fileName, contents)

	reader, err := NewTarGzFileReader(tgz, fileName)
	if err != nil {
		t.Fatalf("cannot create file reader: %v", err)
	}
	defer reader.Close()

	got, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("cannot read file: %v", err)
	}

	if got, want := got, contents; bytes.Compare(got, want) != 0 {
		t.Errorf("expected test, got: %s", string(got))
	}
}

func TestTarGzFileReader_NotFound(t *testing.T) {
	fileName := "test.txt"
	contents := []byte("test")
	tgz := createTarGz(t, fileName, contents)

	reader, err := NewTarGzFileReader(tgz, "not_test.txt")
	if err != nil {
		t.Fatalf("cannot create file reader: %v", err)
	}
	defer reader.Close()

	_, err = ioutil.ReadAll(reader)
	if err != ErrFileNotFound {
		t.Errorf("expected ErrFileNotFound, received: %v", err)
	}
}

func ExampleNewTarGzFileReader() {
	tgz, _ := os.Open("testdata/test.tar.gz")

	reader, _ := NewTarGzFileReader(tgz, "test.txt")
	contents, _ := ioutil.ReadAll(reader)
	fmt.Println(string(contents))
	// Output: test
}
