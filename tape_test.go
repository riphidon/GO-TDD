package poker

import (
	"io/ioutil"
	"testing"
)

func TestTApe_Write(t *testing.T) {
	file, clean := CreateTempFile(t, "12345")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)

	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := ("abc")

	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}
