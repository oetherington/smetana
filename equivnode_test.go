package smetana

import "testing"

func TestRenderEquiv(t *testing.T) {
	node := Equiv("foo", "bar")
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<meta content=\"bar\" http-equiv=\"foo\">", result)
}

func TestRenderRefresh(t *testing.T) {
	node := Refresh(30)
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<meta content=\"30\" http-equiv=\"refresh\">", result)
}

func TestRenderXUaCompatible(t *testing.T) {
	node := XUaCompatible("ie=edge")
	result := RenderHtmlOpts(node, true, nil)
	expected := "<meta content=\"ie=edge\" http-equiv=\"x-ua-compatible\">"
	assertEqual(t, expected, result)
}
