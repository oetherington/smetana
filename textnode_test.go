package smetana

import "testing"

func TestRenderText(t *testing.T) {
	node := Text("Hello world")
	result := RenderOpts(node, true)
	assertEqual(t, "Hello world", result)
}
