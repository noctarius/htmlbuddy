package sanitizer

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"path/filepath"
	"strings"
	"unicode"
)

var AvailableSanitizers = map[string]interface{}{
	"DeleteNodeAndChildren":                DeleteNodeAndChildren,
	"DeleteElementAndMoveChildrenToParent": DeleteElementAndMoveChildrenToParent,
	"AddStyleDeclaration":                  AddStyleDeclaration,
	"ReplaceElementAndReassignChildren":    ReplaceElementAndReassignChildren,
	"SelectParent":                         SelectParent,
	"And":                                  And,
	"AddAttribute":                         AddAttribute,
	"InjectOuterElement":                   InjectOuterElement,
	"AddAttributeWithFunction":             AddAttributeWithFunction,
	"DeleteAttribute":                      DeleteAttribute,
	"Filter":                               Filter,
	"DeleteStyleDeclaration":               DeleteStyleDeclaration,
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
		value := function(node)
		SetAttribute(node, attribute, value)
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
