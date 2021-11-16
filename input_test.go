package golibav_test

import (
	"net/http"
	"runtime"
	"testing"

	av "github.com/demonoid81/golibav"
	"github.com/ssttevee/go-fmterrors"
)

func TestOpenInputReader(t *testing.T) {
	res, err := http.Get("https://archive.org/download/BigBuckBunny_124/Content/big_buck_bunny_720p_surround.mp4")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if _, err = av.OpenInputReader(res.Body); err != nil {
		t.Fatal(err)
	}

	// invoke gc to test finalizer
	runtime.GC()
}

func TestOpenInputWithOpener(t *testing.T) {
	// if _, err := av.OpenInputWithOpener(av.FileOpener, "README.md"); err != nil {
	if _, err := av.OpenInputWithOpener(av.FileOpener, "big_buck_bunny_720p_surround.mp4"); err != nil {
		t.Fatal(fmterrors.FormatString(err))
	}

	// invoke gc to test finalizer
	runtime.GC()
}
