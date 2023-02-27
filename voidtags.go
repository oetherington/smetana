package smetana

import "sort"

// This is a list of all the "void" (or "singleton") tags in valid HTML. It is
// used in a binary search so it must be kept in alphabetical order.
// https://html.spec.whatwg.org/multipage/syntax.html#syntax-tags
// TODO: The spec above allows foreign tags (ie; from the SVG and MathML)
// namespaces to also be void - handle that here
var voidTags = [...]string{
	"area",
	"base",
	"br",
	"col",
	"embed",
	"hr",
	"img",
	"input",
	"link",
	"meta",
	"source",
	"track",
	"wbr",
}

func isVoidTag(tag string) bool {
	index := sort.SearchStrings(voidTags[:], tag)
	return voidTags[index] == tag
}
