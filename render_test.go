package smetana

import "testing"

func TestRenderNodeWithDefaultOptions(t *testing.T) {
	result := RenderHtml(Text("Hello world"))
	assertEqual(t, "Hello world", result)
}
