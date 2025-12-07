package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile(path.Join("subdir", "file1"))

	err = os.MkdirAll(path.Join("subdir", "parent", "child"), 0755)
	check(err)

	createEmptyFile(path.Join("subdir", "parent", "file2"))
	createEmptyFile(path.Join("subdir", "parent", "file3"))
	createEmptyFile(path.Join("subdir", "parent", "child", "file2"))

	c, err := os.ReadDir(path.Join("subdir", "parent"))
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir(path.Join("subdir", "parent", "child"))
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
