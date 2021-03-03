package reader

import (
	"fmt"
	"io"
	"log"
	"os"

	"path/filepath"
	"testing"
)

func TestPortReader(t *testing.T) {
	reader := NewPortReader()
	filename, erFile := copyTestFile()
	if erFile != nil {
		t.Fatal(erFile)
	}
	if er := reader.Init(filename); er != nil {
		t.Fatal(er)
	}

	_, erp := reader.NextPort()
	if erp != nil {
		if erp != io.EOF {
			t.Fatal(erp)
		}
	}
	reader.Close()
	_, err := os.Stat(filename)
	if err == nil {
		t.Error("JSON test file still exists.")
	}
}

func copyTestFile() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	sep := string(os.PathSeparator)
	baseDir := filepath.Dir(fmt.Sprintf("%s%s%s", path, sep, "../"))
	dst := fmt.Sprintf("%s%s%s%s%s", baseDir, sep, "tmp", sep, "ports.json")
	src := fmt.Sprintf("%s%s%s%s%s", baseDir, sep, "test", sep, "ports.json")
	log.Printf("Path of test file: %s\n", dst)

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return "", err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return "", fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		return "", err
	}
	return dst, nil
}
