package sanitizer

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"path/filepath"
	"reflect"
	"strings"
	"unicode"
)

var AvailableSanitizers = map[string]SanitizerFactory{
	"DeleteNodeAndChildren": func(arguments ...string) Sanitizer {
		return DeleteNodeAndChildren
	},
	"DeleteElementAndMoveChildrenToParent": func(arguments ...string) Sanitizer {
		return DeleteElementAndMoveChildrenToParent
	},
	"AddStyleDeclaration": func(arguments ...string) Sanitizer {
		return generateSanitizer(AddStyleDeclaration, 2, arguments)
	},
	"ReplaceElementAndReassignChildren": func(arguments ...string) Sanitizer {
		return generateSanitizer(ReplaceElementAndReassignChildren, 1, arguments)
	},
	"SelectParent": func(arguments ...string) Sanitizer {
		return generateSanitizer(SelectParent, 1, arguments)
	},
}

func generateSanitizer(factory interface{}, expectedArguments int, arguments []string) Sanitizer {
	val := reflect.ValueOf(factory)
	if len(arguments) != 1 {
		panic(fmt.Sprintf("cannot create %s, illegal number of arguments, expected %d", val.Type().Name(), expectedArguments))
	}

	args := transformArguments(arguments)
	ret := reflect.ValueOf(SelectParent).Call(args)
	return ret[0].Interface().(Sanitizer)
}

func transformArguments(arguments []string) []reflect.Value {
	args := make([]reflect.Value, len(arguments))
	for i, arg := range arguments {
		args[i] = reflect.ValueOf(arg)
	}
	return args
}

var Sanitizers = map[string]Sanitizer{
	"head": DeleteNodeAndChildren,

	"body": DeleteElementAndMoveChildrenToParent,

	"html": DeleteElementAndMoveChildrenToParent,

	"a[id^='post-']": DeleteElementAndMoveChildrenToParent,

	"td p": DeleteElementAndMoveChildrenToParent,

	"table": AddStyleDeclaration("width", "100%"),

	"h1,h2,h3,h5,h6,h7,h8,h9": ReplaceElementAndReassignChildren("h4"),

	"p h1, p h2, p h3, p h4, p h5, p h6, p h7, p h8, p h9": SelectParent(DeleteElementAndMoveChildrenToParent),

	"a:not([data-rel='lightbox'])": And(
		AddAttribute("target", "_blank"),
		AddAttribute("rel", "noopener"),
	),

	"img": And(
		InjectOuterElement("a"),
		SelectParent(
			And(
				AddAttribute("data-rel", "lightbox"),
				AddAttributeWithFunction("href", extractHref),
			),
		),
		AddStyleDeclaration("border", "1px solid DarkSlateGray"),
		AddAttributeWithFunction("alt", generateAlt),
		DeleteAttribute("class"),
	),

	"p": Filter(
		isCodeBlock,
		And(
			ReplaceElementAndReassignChildren("pre"),
			AddAttribute("data-rel", ""),
		),
	),
}

func generateAlt(node *html.Node) string {
	if src, ok := GetAttribute(node, "src"); ok {
		href := src.Val

		index := strings.LastIndex(href, "/") + 1
		alt := href[index:]
		alt = strings.TrimSuffix(alt, filepath.Ext(alt))
		alt = strings.Replace(alt, "-", " ", -1)
		alt = strings.Replace(alt, ".", " ", -1)

		builder := new(strings.Builder)
		for i, rune := range alt {
			lastRune := int32(' ')
			if i > 0 {
				lastRune = int32(alt[i-1])
			}
			if lastRune == ' ' && rune != ' ' {
				lastRune = rune
				rune = unicode.ToUpper(rune)
			}
			builder.WriteRune(rune)
		}
		return builder.String()
	}
	return ""
}

func extractHref(node *html.Node) string {
	if src, ok := GetAttribute(node.FirstChild, "src"); ok {
		return src.Val
	}
	return ""
}

func isCodeBlock(node *html.Node) bool {
	if !IsTextOnly(node) {
		return false
	}

	textNode := node.FirstChild
	text := strings.TrimSpace(textNode.Data)
	if !strings.HasPrefix(text, "```") {
		return false
	}

	textNode = node.LastChild
	text = strings.TrimSpace(textNode.Data)
	if !strings.HasSuffix(text, "```") {
		return false
	}
	return true
}

func Filter(filter func(node *html.Node) bool, sanitizer Sanitizer) Sanitizer {
	return func(node, parent *html.Node) error {
		if filter(node) {
			return sanitizer(node, parent)
		}
		return nil
	}
}

func SelectParent(sanitizer Sanitizer) Sanitizer {
	return func(node, parent *html.Node) error {
		node = node.Parent
		return sanitizer(node, node.Parent)
	}
}

func And(sanitizers ... Sanitizer) Sanitizer {
	return func(node, parent *html.Node) error {
		for _, sanitizer := range sanitizers {
			if err := sanitizer(node, parent); err != nil {
				return err
			}
		}
		return nil
	}
}

func DeleteNodeAndChildren(node, parent *html.Node) error {
	RemoveNode(parent, node)
	return nil
}

func DeleteElementAndMoveChildrenToParent(node, parent *html.Node) error {
	children := Children(node)
	for child := range children {
		MoveNode(child, node, parent)
	}
	RemoveNode(parent, node)
	return nil
}

func AddAttributeWithFunction(attribute string, function func(node *html.Node) string) Sanitizer {
	return func(node, parent *html.Node) error {
		SetAttribute(node, attribute, function(node))
		return nil
	}
}

func AddAttribute(attribute, value string) Sanitizer {
	return func(node, parent *html.Node) error {
		SetAttribute(node, attribute, value)
		return nil
	}
}

func DeleteAttribute(attribute string) Sanitizer {
	return func(node, parent *html.Node) error {
		RemoveAttribute(node, attribute)
		return nil
	}
}

func InjectOuterElement(tagName string) Sanitizer {
	return func(node, parent *html.Node) error {
		newParent := CreateTag(tagName)
		parent.InsertBefore(newParent, node)
		MoveNode(node, parent, newParent)
		return nil
	}
}

func AddStyleDeclaration(property, value string) Sanitizer {
	return func(node, parent *html.Node) error {
		style := ParseStyle(node)
		style.SetDeclaration(property, value)
		style.AttachStyle(node)
		return nil
	}
}

func DeleteStyleDeclaration(property string) Sanitizer {
	return func(node, parent *html.Node) error {
		style := ParseStyle(node)
		style.RemoveDeclaration(property)
		style.AttachStyle(node)
		return nil
	}
}

func ReplaceElementAndReassignChildren(newTagName string) Sanitizer {
	return func(node, parent *html.Node) error {
		node.Data = newTagName
		node.DataAtom = atom.Lookup([]byte(newTagName))
		return nil
	}
}
