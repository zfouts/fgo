package command

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// TestPrepareDirectory tests prepareDirectory within the following three cases;
// 1) creating non existing directory, 2) using existing directory,
// 3) trying to make a directory of which name is colliding another file.
func TestPrepareDirectory(t *testing.T) {

	var err error

	// Move temporary directory.
	cd, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
		return
	}
	if err = os.Chdir(os.TempDir()); err != nil {
		t.Error(err.Error())
		return
	}
	defer os.Chdir(cd)

	t.Log("Test with an existing directory.")
	target, err := ioutil.TempDir("", "fgo-test")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer os.RemoveAll(target)

	if err = prepareDirectory(target); err != nil {
		t.Error(err.Error())
	}

	t.Log("Test with an existing file.")
	fp, err := ioutil.TempFile("", "fgo-test2")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fp.Close()

	if prepareDirectory(fp.Name()) == nil {
		t.Error("No error occues when preparing directory with an existing file.")
	}

	if err = os.Remove(fp.Name()); err != nil {
		t.Error(err.Error())
		return
	}

	t.Log("Test without any collisions.")
	if err = prepareDirectory(fp.Name()); err != nil {
		t.Error(err.Error())
	}

}

// TestMakefile tests generated Makefile contains a name given as a parameter.
func TestMakefile(t *testing.T) {

	param := MakefileParam{
		Dest: "test",
	}

	data, err := makefile(&param)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !strings.Contains(string(data), "-d=test") {
		t.Errorf("Generated Makefile was wrong.\n%s", string(data))
	}

}

func TestBrewfile(t *testing.T) {

	param := TemplateRbParam{
		Package:  "test",
		UserName: "abcde",
	}

	data, err := brewfile(&param)
	if err != nil {
		t.Error(err.Error())
		return
	}

	res := string(data)
	if !strings.Contains(res, "class Test") {
		t.Error("Generated file has wrong class name.")
	}
	if !strings.Contains(res, "https://github.com/abcde/test") {
		t.Error("Generated file has wrong URL.")
	}
	if !strings.Contains(res, "bin.install \"test\"") {
		t.Error("Generated file has wrong install command.")
	}
	t.Log(res)

}
