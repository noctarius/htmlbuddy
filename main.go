package main

import (
	"github.com/Joker/hpp"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"instanawordpress/sanitizer"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	inputFile := os.Args[1]
	//outputFile := os.Args[1]

	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	in, err := ioutil.ReadAll(input)
	node, err := html.Parse(strings.NewReader(string(in)))
	if err != nil {
		panic(err)
	}

	for query, sanitizer := range sanitizer.Sanitizers {
		q, err := cascadia.Compile(query)
		if err != nil {
			panic(err)
		}

		nodes := q.MatchAll(node)
		for _, child := range nodes {
			parent := child.Parent

			err := sanitizer(child, parent)
			if err != nil {
				panic(err)
			}
		}
	}

	// Remove outer html
	nodes := sanitizer.Children(node)
	buffer := new(strings.Builder)
	for node := range nodes {
		html.Render(buffer, node)
	}

	println(hpp.PrPrint(buffer.String()))
}
