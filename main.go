package main

import (
	"github.com/Joker/hpp"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"wordpress-sanitizer/sanitizer"
	"io/ioutil"
	"os"
	"strings"
	"flag"
	"github.com/dop251/goja"
	"log"
)

var configFlag = flag.String("configuration", "", "")

func main() {
	flag.Parse()

	config := *configFlag
	runtime := goja.New()

	f, err := os.Open(config)
	if err != nil {
		panic(err)
	}
	script, err := ioutil.ReadAll(f)

	inputFile := flag.Arg(0)

	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	in, err := ioutil.ReadAll(input)
	node, err := html.Parse(strings.NewReader(string(in)))
	if err != nil {
		panic(err)
	}

	registerNativeImplementations(node, runtime)

	_, err = runtime.RunString(string(script))
	if err != nil {
		panic(err)
	}

	// Remove outer html
	nodes := sanitizer.Children(node)
	buffer := new(strings.Builder)
	for node := range nodes {
		html.Render(buffer, node)
	}

	os.Stdout.WriteString(hpp.PrPrint(buffer.String()))
}
func registerNativeImplementations(node *html.Node, runtime *goja.Runtime) {
	global := runtime.GlobalObject()

	for key, value := range sanitizer.AvailableSanitizers {
		global.Set(key, runtime.ToValue(value))
	}

	global.Set("sanitizers", sanitizers(node))
	global.Set("sanitize", sanitize)

	api := runtime.NewObject()
	api.Set("isTextOnly", sanitizer.IsTextOnly)
	api.Set("parseStyle", sanitizer.ParseStyle)
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
	global.Set("api", api)

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
