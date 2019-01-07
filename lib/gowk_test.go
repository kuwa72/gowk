package lib

import (
	"testing"
)

func TestBuildImports(t *testing.T) {
	if "" != buildImports() {
		t.Error("empty arg for not empty")
	}

	if "\"a\";" != buildImports("a") {
		t.Error("a a is not a")
	}

	if "\"a\";\"b\";" != buildImports("a", "b") {
		t.Error("a b is not a b")
	}
}

func TestCreateFileToTempDir(t *testing.T) {
	fn, err := createFileToTempDir("foo")
	if err != nil {
		t.Error(err)
	}
	if "" == fn {
		t.Error("Can not create tempporary directory")
	}
}

func TestRun(t *testing.T) {
	if err := Run("", "", "", "", false, false); err != nil {
		t.Error(err)
	}
	if err := Run("", "", "", "", true, false); err != nil {
		t.Error(err)
	}
}
