package sanitizer

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var NativeAPI = map[string]interface{}{
	"DeleteNodeAndChildren":                DeleteNodeAndChildren,
	"DeleteElementAndMoveChildrenToParent": DeleteElementAndMoveChildrenToParent,
	"SetStyleDeclaration":                  SetStyleDeclaration,
	"ReplaceElementAndReassignChildren":    ReplaceElementAndReassignChildren,
	"SelectParent":                         SelectParent,
	"And":                                  And,
	"SetAttribute":                         SetAttribute0,
	"InjectOuterElement":                   InjectOuterElement,
	"SetAttributeWithExtractor":            SetAttributeWithExtractor,
	"DeleteAttribute":                      DeleteAttribute,
	"Filter":                               Filter,
	"Filters":                              Filters,
	"DeleteStyleDeclaration":               DeleteStyleDeclaration,
}

func Filter(predicate Predicate, sanitizer Sanitizer) Sanitizer {
	return func(node, parent *html.Node) error {
		if predicate(node) {
			return sanitizer(node, parent)
		}
		return nil
	}
}

func Filters(predicates ...Predicate) Predicate {
	return func(node *html.Node) bool {
		for _, predicate := range predicates {
			if !predicate(node) {
				return false
			}
		}
		return true
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

func SetAttributeWithExtractor(attribute string, extractor Extractor) Sanitizer {
	return func(node, parent *html.Node) error {
		value := extractor(node)
		SetAttribute(node, attribute, value)
		return nil
	}
}

func SetAttribute0(attribute, value string) Sanitizer {
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

func SetStyleDeclaration(property, value string) Sanitizer {
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
