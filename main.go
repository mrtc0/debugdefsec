package main

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/aquasecurity/defsec/parsers/dockerfile/parser"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <parser> <file>\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dockerfile":
		p := parser.New()
		dockerfile, err := p.ParseFile(os.Args[2])
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		json, err := json.Marshal(&dockerfile)
		if err != nil {
			fmt.Printf("Marshal failed: %s", err)
		}
		fmt.Printf("%s\n", json)
	}

	os.Exit(0)
}
