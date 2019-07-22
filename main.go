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

		// build plugin
		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", soFile, goFile)
		if err := cmd.Run(); err != nil {
			panic("failed to compile plugin: " + err.Error())
		}

		p, err := plugin.Open(soFile)
		if err != nil {
			panic("cannot open build file: " + err.Error())
		}

		f, err := p.Lookup("Iteration")
		if err != nil {
			panic("cannot lookup Iteration(): " + err.Error())
		}

		// function `Iteration` should prints respective the value of `i`
		// but will always print first value in the file: `Iteration: 0`
		f.(func())()

		if err := os.Remove(soFile); err != nil {
			panic("cannot remove file: " + err.Error())
		}
	}
}
