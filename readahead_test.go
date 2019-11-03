package readahead

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadahead(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	testData := []byte("hello world!\n")
	for i := 0; i < 2<<14; i++ {
		if _, err := tmpfile.Write(testData); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	// These pages should be in memory anyways.
	if err := Readahead(int(tmpfile.Fd()), 0, 2<<10); err != nil {
		t.Fatal(err)
	}
}
