package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

const (
	// ExitOk will be returned whenever application runs without any erros
	ExitOk int = iota
	// ExitError will be returned whenever application interrupted by occurred error
	ExitError
)

/*
type Section struct {
	about string
	questions []Question
}
*/

type Question struct {
	Title       string `toml:"title"`
	Purpose     string `toml:"purpose"`
	Description string `toml:"description"`
}

type Section struct {
	About     string     `toml:"about"`
	Questions []Question `toml:"question"`
}

type Challenge struct {
	Sections []Section `toml:"section"`
}

var text = `
[[section]]
about = "I'm describing about this section"
[[section.question]]
title = "title is shown as headline text to applicant"
purpose = "purpose is an internal note about 'what can be seen through this question'"
description = "Description will be an question text that they will be answered"
[[section.question]]
title = "the next"
purpose = "the next"
description = "the next"
`

func main() {
	os.Exit(Run(os.Args))
}

func Run(args []string) int {
	loaded, err := toml.Load(text)
	if err != nil {
		println(err.Error())
		return ExitError
	}
	println("----- string -----")
	println(loaded.String())
	challenge := Challenge{}
	toml.Unmarshal([]byte(text), &challenge)

	println("===== Unmarshal =====")
	fmt.Printf("Value: %v\n", challenge)

	return ExitOk
}
