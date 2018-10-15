package main

import (
	"bufio"
	"errors"
	"flag"
	"github.com/Joker/hpp"
	"github.com/andybalholm/cascadia"
	"github.com/dop251/goja"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"wordpress-sanitizer/sanitizer"
)

var configFlag = flag.String("configuration", "", "")

func main() {
	flag.Parse()

	config := *configFlag
	runtime := goja.New()

	script, err := readFile(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	var input string
	inputFile := flag.Arg(0)
	if inputFile == "-" {
		reader := bufio.NewReader(os.Stdin)
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err.Error())
		}
		input = string(data)

	} else {
		input, err = readFile(inputFile)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	node, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err.Error())
	}

	registerNativeImplementations(node, runtime)

	_, err = runtime.RunString(string(script))
	if err != nil {
		log.Fatal(err.Error())
	}

	// Remove outer html
	nodes := sanitizer.Children(node)
	buffer := new(strings.Builder)
	for node := range nodes {
		html.Render(buffer, node)
	}

	pretty := hpp.PrPrint(buffer.String())
	pretty = strings.Replace(pretty, "$$", "&", -1)

	_, err = os.Stdout.WriteString(pretty)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func readFile(file string) (string, error) {
	url := strings.ToLower(file)
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return downloadFile(file)
	}

	var input io.Reader
	if file == "-" {
		input = bufio.NewReader(os.Stdin)

	} else {
		var err error
		input, err = os.Open(file)
		if err != nil {
			return "", err
		}
		defer input.(io.ReadCloser).Close()
	}

	data, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func downloadFile(file string) (string, error) {
	response, err := http.Get(file)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	builder := new(strings.Builder)
	_, err = io.Copy(builder, response.Body)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

func registerNativeImplementations(node *html.Node, runtime *goja.Runtime) {
	global := runtime.GlobalObject()

	for key, value := range sanitizer.NativeAPI {
		global.Set(key, runtime.ToValue(value))
	}

	global.Set("sanitizers", sanitizers(node))
	global.Set("sanitize", sanitize)

	api := runtime.NewObject()
	api.Set("isTextOnly", sanitizer.IsTextOnly)
	api.Set("parseStyle", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 1 {
			return runtime.NewGoError(errors.New("parseStyle requires 1 arguments: Node"))
		}

		argument := call.Argument(0)
		node, ok := argument.Export().(*html.Node)
		if !ok {
			return runtime.NewGoError(errors.New("first argument of parseStyle cannot be converted to a Node"))
		}

		style := sanitizer.ParseStyle(node)

		s := runtime.NewObject()
		s.Set("getDeclaration", func(key string) (string, bool) {
			return style.Declaration(key)
		})
		s.Set("setDeclaration", func(key, value string) {
			style.SetDeclaration(key, value)
		})
		s.Set("removeDeclaration", func(key string) {
			style.RemoveDeclaration(key)
		})
		s.Set("computeStyle", func() string {
			return style.ComputeStyle()
		})
		s.Set("attachStyle", func(node *html.Node) {
			style.AttachStyle(node)
		})
		return s
	})
	api.Set("hasAttribute", sanitizer.HasAttribute)
	api.Set("getAttribute", sanitizer.GetAttribute)
	api.Set("setAttribute", sanitizer.SetAttribute)
	api.Set("removeAttribute", sanitizer.RemoveAttribute)
	api.Set("createTag", sanitizer.CreateTag)
	api.Set("moveNode", sanitizer.MoveNode)
	api.Set("removeNode", sanitizer.RemoveNode)
	api.Set("emptyNode", sanitizer.EmptyNode)
	api.Set("replaceNode", sanitizer.ReplaceNode)
	api.Set("children", sanitizer.Children)
	api.Set("cloneNode", sanitizer.CloneNode)
	api.Set("createTextNode", sanitizer.CreateTextNode)
	api.Set("appendNode", sanitizer.AppendNode)
	global.Set("api", api)

	buildNodeType(runtime)
	buildAtom(runtime)

	console := runtime.NewObject()
	console.Set("log", func(msg string) {
		log.Println(msg)
	})
	global.Set("console", console)

	s := global.Get("String")
	p := s.ToObject(runtime).Get("prototype")
	po := p.ToObject(runtime)
	po.Set("startsWith", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(strings.HasPrefix(call.This.String(), call.Argument(0).String()))
	})
	po.Set("endsWith", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(strings.HasSuffix(call.This.String(), call.Argument(0).String()))
	})
	po.Set("replaceAll", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(strings.Replace(call.This.String(), call.Argument(0).String(), call.Argument(1).String(), -1))
	})
}

func sanitize(selector string, sanitizer sanitizer.Sanitizer) *configuration {
	return &configuration{selector: selector, sanitizer: sanitizer}
}

func sanitizers(node *html.Node) func(...*configuration) {
	return func(sanitizers ... *configuration) {
		for _, s := range sanitizers {
			q, err := cascadia.Compile(s.selector)
			if err != nil {
				panic(err)
			}

			nodes := q.MatchAll(node)
			for _, child := range nodes {
				parent := child.Parent

				err := s.sanitizer(child, parent)
				if err != nil {
					panic(err)
				}
			}
		}

	}
}

type configuration struct {
	selector  string
	sanitizer sanitizer.Sanitizer
}
