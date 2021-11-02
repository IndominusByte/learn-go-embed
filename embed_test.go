package learngoembed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed default.jpg
var image []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("default_new.jpg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files_1/* files_2/*
var content embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := content.ReadFile("files_1/a.txt")
	fmt.Println(string(a))

	b, _ := content.ReadFile("files_2/b.txt")
	fmt.Println(string(b))
}

//go:embed files_1/*.txt
var content2 embed.FS

func TestPathMatch(t *testing.T) {
	dirEntries, _ := content2.ReadDir("files_1")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			cont, _ := content.ReadFile("files_1/" + entry.Name())
			fmt.Println(entry.Name(), ":", string(cont))
		}
	}
}
