package smetana

import "testing"

func TestRenderFragmentNode(t *testing.T) {
	node := Fragment(Span("Foo"), Div("Bar"))
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<span>Foo</span><div>Bar</div>", result)
}

func TestRenderEmptyFragmentNode(t *testing.T) {
	node := Fragment()
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "", result)
}

func TestCanAppendChildrenToEmptyFragment(t *testing.T) {
	node := Fragment()
	node.AssignChildren([]Node{Div("Bar")})
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<div>Bar</div>", result)
}

func TestCanAppendChildrenToNonEmptyFragment(t *testing.T) {
	node := Fragment(Span("Foo"))
	node.AssignChildren([]Node{Div("Bar")})
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<span>Foo</span><div>Bar</div>", result)
}
