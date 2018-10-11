package sanitizer

import (
	"github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

type Sanitizer = func(node, parent *html.Node) error

type Extractor = func(node *html.Node) string

type Predicate = func(node *html.Node) bool

type Style struct {
	styles []*css.Declaration
}

func (s *Style) Declaration(key string) (string, bool) {
	for _, style := range s.styles {
		if strings.EqualFold(style.Property, key) {
			return style.Value, true
		}
	}

	return "", false
}

func (s *Style) SetDeclaration(key, value string) {
	for _, style := range s.styles {
		if strings.EqualFold(style.Property, key) {
			style.Value = value
		}
	}

	s.styles = append(s.styles, &css.Declaration{Property: key, Value: value, Important: false})
}

func (s *Style) RemoveDeclaration(key string) {
	newStyles := make([]*css.Declaration, 0)
	for _, style := range s.styles {
		if !strings.EqualFold(style.Property, key) {
			newStyles = append(newStyles, style)
		}
	}
	s.styles = newStyles
}

func (s *Style) ComputeStyle() string {
	builder := new(strings.Builder)
	for _, style := range s.styles {
		builder.WriteString(style.StringWithImportant(false))
	}
	return builder.String()
}

func (s *Style) AttachStyle(node *html.Node) {
	SetAttribute(node, "style", s.ComputeStyle())
}

func CreateTextNode(value string) *html.Node {
	node := new(html.Node)
	node.Type = html.TextNode
	node.Data = value
	node.DataAtom = 0
	node.Attr = nil
	return node
}

func IsTextOnly(node *html.Node) bool {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type != html.TextNode {
			return false
		}
	}
	return true
}

func ParseStyle(node *html.Node) *Style {
	attribute, ok := GetAttribute(node, "style")
	if ok {
		styles, err := parser.ParseDeclarations(attribute.Val)
		if err != nil {
			panic(err)
		}
		return &Style{styles: styles}
	}
	return &Style{styles: make([]*css.Declaration, 0)}
}

func HasAttribute(node *html.Node, attribute string) bool {
	for _, attr := range node.Attr {
		if strings.EqualFold(attr.Key, attribute) {
			return true
		}
	}
	return false
}

func GetAttribute(node *html.Node, attribute string) (html.Attribute, bool) {
	for _, attr := range node.Attr {
		if strings.EqualFold(attr.Key, attribute) {
			return attr, true
		}
	}
	return html.Attribute{}, false
}

func SetAttribute(node *html.Node, attribute, value string) {
	for _, attr := range node.Attr {
		if strings.EqualFold(attr.Key, attribute) {
			attr.Val = value
			return
		}
	}
	node.Attr = append(node.Attr, html.Attribute{Key: attribute, Val: value})
}

func RemoveAttribute(node *html.Node, attribute string) {
	attributes := make([]html.Attribute, 0)
	for _, attr := range node.Attr {
		if !strings.EqualFold(attr.Key, attribute) {
			attributes = append(attributes, attr)
		}
	}
	node.Attr = attributes
}

func CreateTag(tag string) *html.Node {
	a := atom.Lookup([]byte(tag))

	node := new(html.Node)
	node.Type = html.ElementNode
	node.Data = tag
	node.DataAtom = a
	node.Attr = make([]html.Attribute, 0)

	return node
}

func AppendNode(node, parent *html.Node) {
	parent.AppendChild(node)
}

func MoveNode(node, oldParent, newParent *html.Node) {
	oldParent.RemoveChild(node)
	newParent.AppendChild(node)
}

func RemoveNode(parent, node *html.Node) {
	EmptyNode(node)
	parent.RemoveChild(node)
}

func EmptyNode(node *html.Node) {
	for c := node.FirstChild; c != nil; c = node.FirstChild {
		node.RemoveChild(c)
	}
}

func ReplaceNode(parent, newNode, oldNode *html.Node) {
	parent.InsertBefore(newNode, oldNode)
	parent.RemoveChild(oldNode)
}

func Children(parent *html.Node) <-chan *html.Node {
	next := make(chan *html.Node)
	go func() {
		for child := parent.FirstChild; child != nil; {
			currentChild := child
			child = child.NextSibling
			next <- currentChild
		}
		close(next)
	}()
	return next
}

func CloneNode(n *html.Node) *html.Node {
	nn := &html.Node{
		Type:     n.Type,
		DataAtom: n.DataAtom,
		Data:     n.Data,
		Attr:     make([]html.Attribute, len(n.Attr)),
	}

	copy(nn.Attr, n.Attr)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nn.AppendChild(CloneNode(c))
	}

	return nn
}
