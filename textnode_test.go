package smetana

import "testing"

func TestRenderText(t *testing.T) {
	node := Text("Hello world")
	result := Render(node)
	assertEqual(t, "Hello world", result)
}
