sanitizers(
    sanitize("head", DeleteNodeAndChildren),
    sanitize("body", DeleteElementAndMoveChildrenToParent),
    sanitize("html", DeleteElementAndMoveChildrenToParent),
    sanitize("a[id^='post-']", DeleteElementAndMoveChildrenToParent),
    sanitize("td p", DeleteElementAndMoveChildrenToParent),
    sanitize("table", SetStyleDeclaration("width", "100%")),
    sanitize("h1, h2, h3, h5, h6, h7, h8, h9", ReplaceElementAndReassignChildren("h4")),
    sanitize("p h1, p h2, p h3, p h4, p h5, p h6, p h7, p h8, p h9", SelectParent(
        DeleteElementAndMoveChildrenToParent)
    ),
    sanitize("a:not([data-rel='lightbox'])", And(
        SetAttribute("target", "_blank"),
        SetAttribute("rel", "noopener")
    )),
    sanitize("img", And(
        InjectOuterElement("a"),
        SelectParent(
            And(
                SetAttribute("data-rel", "lightbox"),
                SetAttributeWithExtractor("href", extractHref)
            )
        ),
        SetStyleDeclaration("border", "1px solid DarkSlateGray"),
        SetAttributeWithExtractor("alt", generateAlt),
        DeleteAttribute("class")
    )),
    sanitize("p", Filter(
        isCodeBlock,
        And(
            ReplaceElementAndReassignChildren("code"),
            SetAttribute("class", "java"),
            InjectOuterElement("pre"),
            fixCodeBlock
        )
    ))
);

function fixCodeBlock(node) {
    var first = node.FirstChild;
    var last = node.LastChild;

    first.Data = first.Data.substring(3);
    last.Data = last.Data.substring(0, last.Data.length - 4);
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

function isCodeBlock(node) {
    if (!api.isTextOnly(node)) {
        return false;
    }

    var textNode = node.FirstChild;
    var text = textNode.Data.trim();

    if (!text.startsWith("```")) {
        return false;
    }

    textNode = node.LastChild;
    text = textNode.Data.trim();
    return text.endsWith("```");
}
