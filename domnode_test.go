package smetana

import "testing"

func TestRenderDiv(t *testing.T) {
	node := Div(Attrs{"class": "foo"}, Children{Text("bar")})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<div class=\"foo\">bar</div>", result)
}

func TestRenderSpan(t *testing.T) {
	node := Span(Attrs{"class": "foo"}, Children{Text("bar")})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<span class=\"foo\">bar</span>", result)
}

func TestRenderHead(t *testing.T) {
	node := Head(Attrs{"class": "foo"}, Children{Text("bar")})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<head class=\"foo\">bar</head>", result)
}

func TestRenderBody(t *testing.T) {
	node := Body(Attrs{"class": "foo"}, Children{Text("bar")})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<body class=\"foo\">bar</body>", result)
}

func TestRenderTitle(t *testing.T) {
	node := Title("hello world")
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<title>hello world</title>", result)
}

func TestRenderLink(t *testing.T) {
	node := Link("stylesheet", "/main.css")
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<link href=\"/main.css\" rel=\"stylesheet\" />", result)
}

func TestRenderCharset(t *testing.T) {
	node := Charset("ASCII")
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<meta charset=\"ASCII\" />", result)
}

func TestRenderCharsetDefault(t *testing.T) {
	node := Charset("")
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<meta charset=\"UTF-8\" />", result)
}

func TestRenderRefresh(t *testing.T) {
	node := Refresh(30)
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<meta content=\"30\" http-equiv=\"refresh\" />", result)
}

func TestRenderBase(t *testing.T) {
	node := Base("https://example.com/")
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<base href=\"https://example.com/\" target=\"_blank\" />", result)
}

func TestRenderScript(t *testing.T) {
	node := Script("/main.js")
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<script src=\"/main.js\" />", result)
}

func TestRenderInlineScript(t *testing.T) {
	node := InlineScript("alert('foo')")
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<script>alert('foo')</script>", result)
}

func TestRenderStyle(t *testing.T) {
	node := Style("body{background:red}")
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<style>body{background:red}</style>", result)
}
