package smetana

import "testing"

func TestRenderHtmlNode(t *testing.T) {
	node := Html(Attrs{"lang": "en"}, Children{Head(), Body()})
	result := RenderHtmlOpts(node, true, nil)
	assertEqual(
		t,
		"<!DOCTYPE html>\n<html lang=\"en\"><head></head><body></body></html>",
		result,
	)
}
