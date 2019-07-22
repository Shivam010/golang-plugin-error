package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"plugin"
)

const (
	goFile = `example.go`
	soFile = `example.so`

	contentTemplate = `package main

import "fmt"

func Iteration() {
	p := "Iteration: %d"
	fmt.Println(p)
}
`
)

func main() {

	file, err := os.Create(goFile);
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {

		code := fmt.Sprintf(contentTemplate, i)
		// modifying the content of the goFile
		if _, err := file.WriteAt([]byte(code), 0); err != nil {
			panic("cannot write file: " + err.Error())
		}

		// content Check
		str, err := ioutil.ReadFile(goFile)
		if err != nil {
			panic("cannot read file: " + err.Error())
		}

		// just to ensure code is changing
		if string(str) != code {
			panic("code is not same")
		}

		// build plugin with a unique name for soFile
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", string(48+i)+soFile, goFile)
		if err := cmd.Run(); err != nil {
			panic("failed to compile plugin: " + err.Error())
		}

		p, err := plugin.Open(string(48+i)+soFile)
		if err != nil {
			panic("cannot open build file: " + err.Error())
		}

		f, err := p.Lookup("Iteration")
		if err != nil {
			panic("cannot lookup Iteration(): " + err.Error())
		}

		// function `Iteration`, now prints respective the value of `i`
		// as we have explicitly generated a new so File instead of over
		// writing it
		f.(func())()

		if err := os.Remove(string(48+i)+soFile); err != nil {
			panic("cannot remove file: " + err.Error())
		}
	}
}
