package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

//go:embed template
var box embed.FS

func main() {

	validateFileExists := func(input string) error {
		dir := filepath.Base(input)
		_, err := os.Stat(dir)
		if err == nil {
			return fmt.Errorf("%s already exists, remove it first to generate.\n", err)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Admin Go Package",
		AllowEdit: true,
		Default:   "github.com/theplant/myadmin",
	}

	pkg, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	if err = validateFileExists(pkg); err != nil {
		sel := promptui.Select{
			Label: "Directory exists, Overwrite?",
			Items: []string{
				"Yes",
				"No",
			},
		}
		result, _, _ := sel.Run()
		if result != 0 {
			return
		}
	}

	dir := filepath.Base(pkg)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}

	fs.WalkDir(box, "template", func(path string, d fs.DirEntry, err1 error) error {
		if d != nil && d.IsDir() {
			return nil
		}
		newPath := strings.ReplaceAll(path, "template/", "")
		fp := filepath.Join(dir, newPath)
		err := os.MkdirAll(filepath.Dir(fp), 0755)
		if err != nil {
			panic(err)
		}
		var f *os.File
		f, err = os.Create(fp)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		content, err := box.ReadFile(path)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fp, []byte(content), 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println(fp, "generated")
		return err
	})

	fmt.Println("Done")

	replaceInFiles(dir, "github.com/goplaid/x/goplaid/template", pkg)
	replaceInFiles(dir, "GoplaidPackageName", dir)

	if _, err = os.Stat(filepath.Join(dir, "go.mod")); err != nil {
		cmd := exec.Command("go", "mod", "init", pkg)
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("\ncd %s && docker-compose up -d && source dev_env && go run main.go\nto start your project\n", dir)
}

func replaceInFiles(baseDir string, from, to string) {
	var fileList []string
	err := filepath.Walk(baseDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range fileList {
		replaceInFile(file, from, to)
	}
}

func replaceInFile(filepath, from, to string) {
	read, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	newContents := strings.Replace(string(read), from, to, -1)

	// fmt.Println(newContents)

	err = ioutil.WriteFile(filepath, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
