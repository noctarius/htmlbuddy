sanitizers(
    // Remove implicit head, added by html parser
    sanitize("head", DeleteNodeAndChildren),

    // Remove implicit body, added by html parser
    sanitize("body", DeleteElementAndMoveChildrenToParent),

    // Remove implicit head, added by html parser
    sanitize("html", DeleteElementAndMoveChildrenToParent),

    // Remove anchor elements from headings
    sanitize("a[id^='post-']", DeleteElementAndMoveChildrenToParent),

    // Remove paragraph from tables
    sanitize("td p", DeleteElementAndMoveChildrenToParent),

    // Adjust table width
    sanitize("table", And(
        DeleteAttribute("width"),
        SetStyleDeclaration("width", "100%")
    )),

    // Set all headers to h4
    sanitize("h1, h2, h3, h5, h6, h7, h8, h9", ReplaceElementAndReassignChildren("h4")),

    // Remove possible paragraphs before headings
    sanitize("p h1, p h2, p h3, p h4, p h5, p h6, p h7, p h8, p h9", SelectParent(
        DeleteElementAndMoveChildrenToParent
    )),

    // Setup links with lightbox for images
    sanitize("a:not([data-rel='lightbox'])", And(
        SetAttribute("target", "_blank"),
        SetAttribute("rel", "noopener")
    )),

    // Configure images to use lightbox and have a slight border,
    // also generates alternative text from filename
    sanitize("img", And(
        InjectOuterElement("a"),
        SelectParent(
            And(
                SetAttribute("data-rel", "lightbox"),
                SetAttributeWithExtractor("href", extractHref)
            )
        ),
        SetAttributeWithExtractor("alt", generateAlt),
        DeleteAttribute("class"),
        DeleteAttribute("height"),
        SetAttribute("width", "730")
    )),

    // Remove attributes added by google spreadsheet
    sanitize("td", And(
        DeleteAttribute("width"),
        DeleteAttribute("data-sheets-value")
    )),

    // Remove empty links
    sanitize("a", Filter(
        isEmptyNode, DeleteNodeAndChildren
    )),

    // Add nbsp to empty tds
    sanitize("td", Filter(
        isEmptyNode, pushNbsp
    ))
);

function pushNbsp(node) {
    api.appendNode(api.createTextNode("$$nbsp;"), node);
}

function isEmptyNode(node) {
    for (var child = node.FirstChild; child != null; child = child.NextSibling) {
        if (child.Type === NodeType.TextNode) {
            if (child.Data.trim().length === 0) {
                continue;
            }
        }
        return false;
    }
    return true;
}

function extractHref(node) {
    var ret = api.getAttribute(node.FirstChild, "src");
    if (ret && ret[1]) {
        return ret[0].Val;
    }
    return ""
}

function generateAlt(node) {
    var ret = api.getAttribute(node, "src");
    if (ret && ret[1]) {
        var href = ret[0].Val;

        var index = href.lastIndexOf("/") + 1;
        var alt = href.substring(index);

        alt = alt.substring(0, alt.lastIndexOf("."));
        alt = alt.replaceAll("-", " ", -1);
        alt = alt.replaceAll(".", " ", -1);

        var alternative = "";
        for (var i = 0; i < alt.length; i++) {
            var lastRune = ' ';
            if (i > 0) {
                lastRune = alt.charAt(i - 1);
            }

            var rune = alt.charAt(i);
            if (lastRune === ' ' && rune !== ' ') {
                rune = rune.toUpperCase();
            }
            alternative += rune;
        }

        return alternative;
    }

    return ""
}
