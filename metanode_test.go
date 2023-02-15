package smetana

import "testing"

func TestRenderMeta(t *testing.T) {
	node := Meta("foo", "bar")
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<meta content=\"bar\" name=\"foo\" />", result)
}

func TestRenderKeywords(t *testing.T) {
	node := Keywords("some,keywords")
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<meta content=\"some,keywords\" name=\"keywords\" />", result)
}

func TestRenderDescription(t *testing.T) {
	node := Description("a description")
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<meta content=\"a description\" name=\"description\" />", result)
}

func TestRenderAuthor(t *testing.T) {
	node := Author("A. Developer")
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<meta content=\"A. Developer\" name=\"author\" />", result)
}

func TestRenderViewport(t *testing.T) {
	node := Viewport("some-value")
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<meta content=\"some-value\" name=\"viewport\" />", result)
}

func TestRenderViewportDefault(t *testing.T) {
	node := Viewport("")
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(t, "<meta content=\"width=device-width, initial-scale=1.0\" name=\"viewport\" />", result)
}
