package smetana

import "testing"

func TestRenderMeta(t *testing.T) {
	node := Meta("foo", "bar")
	result := Render(node)
	assertEqual(t, "<meta name=\"foo\" content=\"bar\" />", result)
}

func TestRenderKeywords(t *testing.T) {
	node := Keywords("some,keywords")
	result := Render(node)
	assertEqual(t, "<meta name=\"keywords\" content=\"some,keywords\" />", result)
}

func TestRenderDescription(t *testing.T) {
	node := Description("a description")
	result := Render(node)
	assertEqual(t, "<meta name=\"description\" content=\"a description\" />", result)
}

func TestRenderAuthor(t *testing.T) {
	node := Author("A. Developer")
	result := Render(node)
	assertEqual(t, "<meta name=\"author\" content=\"A. Developer\" />", result)
}

func TestRenderViewport(t *testing.T) {
	node := Viewport("some-value")
	result := Render(node)
	assertEqual(t, "<meta name=\"viewport\" content=\"some-value\" />", result)
}

func TestRenderViewportDefault(t *testing.T) {
	node := Viewport("")
	result := Render(node)
	assertEqual(t, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />", result)
}
