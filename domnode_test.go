package smetana

import (
	"log"
	"strings"
	"testing"
)

func TestRenderDomNodeWithAttributes(t *testing.T) {
	node := buildDomNode("div", []any{Attrs{"class": "foo"}})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<div class=\"foo\" />", result)
}

func TestRenderDomNodeWithChildren(t *testing.T) {
	node := buildDomNode("div", []any{Children{Text("bar")}})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<div>bar</div>", result)
}

func TestRenderDomNodeWithSingleAttribute(t *testing.T) {
	node := buildDomNode("div", []any{Attr{"foo", "bar"}})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<div foo=\"bar\" />", result)
}

func TestRenderDomNodeWithSingleChild(t *testing.T) {
	node := buildDomNode("div", []any{Text("bar")})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<div>bar</div>", result)
}

func TestRenderDomNodeWithClasses(t *testing.T) {
	node := buildDomNode("div", []any{Classes{"foo": true, "bar": false}})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<div class=\"foo\" />", result)
}

func TestRenderDomNodeWithClassName(t *testing.T) {
	node := buildDomNode("div", []any{ClassName("foo")})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<div class=\"foo\" />", result)
}

func TestRenderDomNodeWithStrings(t *testing.T) {
	node := buildDomNode("div", []any{Classes{"a": true}})
	result := RenderOpts(node, true, nil)
	assertEqual(t, "<div class=\"a\" />", result)
}

func TestDomNodeReportsErrors(t *testing.T) {
	node := buildDomNode("div", []any{3})
	var target strings.Builder
	logger := log.New(&target, "", 0)
	result := RenderOpts(node, true, logger)
	assertEqual(t, "<div />", result)
	output := strings.TrimSpace(target.String())
	assertEqual(t, "Invalid DomNode argument: 3", output)
}

type DomNodeTest struct {
	node     DomNode
	expected string
}

func TestRenderIndividualDomNodes(t *testing.T) {
	testCases := []DomNodeTest{
		{A(Text("foo")), "<a>foo</a>"},
		{
			AHref("/index.html", Text("Home")),
			"<a href=\"/index.html\">Home</a>",
		},
		{Abbr(Text("foo")), "<abbr>foo</abbr>"},
		{Address(Text("foo")), "<address>foo</address>"},
		{Area(Text("foo")), "<area>foo</area>"},
		{Article(Text("foo")), "<article>foo</article>"},
		{Aside(Text("foo")), "<aside>foo</aside>"},
		{Audio(Text("foo")), "<audio>foo</audio>"},
		{B(Text("foo")), "<b>foo</b>"},
		{Base(Text("foo")), "<base>foo</base>"},
		{
			BaseHref("https://example.com/"),
			"<base href=\"https://example.com/\" target=\"_blank\" />",
		},
		{Bdi(Text("foo")), "<bdi>foo</bdi>"},
		{Bdo(Text("foo")), "<bdo>foo</bdo>"},
		{Blockquote(Text("foo")), "<blockquote>foo</blockquote>"},
		{Br(), "<br />"},
		{Button(Text("foo")), "<button>foo</button>"},
		{Canvas(Text("foo")), "<canvas>foo</canvas>"},
		{Caption(Text("foo")), "<caption>foo</caption>"},
		{
			Body(Attrs{"class": "foo"}, Children{Text("bar")}),
			"<body class=\"foo\">bar</body>",
		},
		{Charset("ASCII"), "<meta charset=\"ASCII\" />"},
		{Charset(""), "<meta charset=\"UTF-8\" />"},
		{Cite(Text("foo")), "<cite>foo</cite>"},
		{Code(Text("foo")), "<code>foo</code>"},
		{Col(Text("foo")), "<col>foo</col>"},
		{Colgroup(Text("foo")), "<colgroup>foo</colgroup>"},
		{Data(Text("foo")), "<data>foo</data>"},
		{Datalist(Text("foo")), "<datalist>foo</datalist>"},
		{Dd(Text("foo")), "<dd>foo</dd>"},
		{Del(Text("foo")), "<del>foo</del>"},
		{Details(Text("foo")), "<details>foo</details>"},
		{Dfn(Text("foo")), "<dfn>foo</dfn>"},
		{Dialog(Text("foo")), "<dialog>foo</dialog>"},
		{
			Div(Attrs{"class": "foo"}, Children{Text("bar")}),
			"<div class=\"foo\">bar</div>",
		},
		{Dl(Text("foo")), "<dl>foo</dl>"},
		{Dt(Text("foo")), "<dt>foo</dt>"},
		{Em(Text("foo")), "<em>foo</em>"},
		{Embed(Text("foo")), "<embed>foo</embed>"},
		{Fieldset(Text("foo")), "<fieldset>foo</fieldset>"},
		{Figcaption(Text("foo")), "<figcaption>foo</figcaption>"},
		{Figure(Text("foo")), "<figure>foo</figure>"},
		{Footer(Text("foo")), "<footer>foo</footer>"},
		{Form(Text("foo")), "<form>foo</form>"},
		{H(1, Text("foo")), "<h1>foo</h1>"},
		{H(2, Text("foo")), "<h2>foo</h2>"},
		{H(3, Text("foo")), "<h3>foo</h3>"},
		{H(4, Text("foo")), "<h4>foo</h4>"},
		{H(5, Text("foo")), "<h5>foo</h5>"},
		{H(6, Text("foo")), "<h6>foo</h6>"},
		{H1(Text("foo")), "<h1>foo</h1>"},
		{H2(Text("foo")), "<h2>foo</h2>"},
		{H3(Text("foo")), "<h3>foo</h3>"},
		{H4(Text("foo")), "<h4>foo</h4>"},
		{H5(Text("foo")), "<h5>foo</h5>"},
		{H6(Text("foo")), "<h6>foo</h6>"},
		{Head(Text("foo")), "<head>foo</head>"},
		{Header(Text("foo")), "<header>foo</header>"},
		{Hr(Text("foo")), "<hr>foo</hr>"},
		{I(Text("foo")), "<i>foo</i>"},
		{Iframe(Text("foo")), "<iframe>foo</iframe>"},
		{Img(Text("foo")), "<img>foo</img>"},
		{Input(Text("foo")), "<input>foo</input>"},
		{Ins(Text("foo")), "<ins>foo</ins>"},
		{Kbd(Text("foo")), "<kbd>foo</kbd>"},
		{Label(Text("foo")), "<label>foo</label>"},
		{Legend(Text("foo")), "<legend>foo</legend>"},
		{Li(Text("foo")), "<li>foo</li>"},
		{Link(Text("foo")), "<link>foo</link>"},
		{
			LinkHref("stylesheet", "/main.css"),
			"<link href=\"/main.css\" rel=\"stylesheet\" />",
		},
		{Main(Text("foo")), "<main>foo</main>"},
		{Map(Text("foo")), "<map>foo</map>"},
		{Mark(Text("foo")), "<mark>foo</mark>"},
		{Meter(Text("foo")), "<meter>foo</meter>"},
		{Nav(Text("foo")), "<nav>foo</nav>"},
		{Noscript(Text("foo")), "<noscript>foo</noscript>"},
		{Object(Text("foo")), "<object>foo</object>"},
		{Ol(Text("foo")), "<ol>foo</ol>"},
		{Optgroup(Text("foo")), "<optgroup>foo</optgroup>"},
		{Option(Text("foo")), "<option>foo</option>"},
		{Output(Text("foo")), "<output>foo</output>"},
		{P(Text("foo")), "<p>foo</p>"},
		{Param(Text("foo")), "<param>foo</param>"},
		{Picture(Text("foo")), "<picture>foo</picture>"},
		{Pre(Text("foo")), "<pre>foo</pre>"},
		{Progress(Text("foo")), "<progress>foo</progress>"},
		{Q(Text("foo")), "<q>foo</q>"},
		{
			Refresh(30),
			"<meta content=\"30\" http-equiv=\"refresh\" />",
		},
		{Rp(Text("foo")), "<rp>foo</rp>"},
		{Rt(Text("foo")), "<rt>foo</rt>"},
		{Ruby(Text("foo")), "<ruby>foo</ruby>"},
		{S(Text("foo")), "<s>foo</s>"},
		{Samp(Text("foo")), "<samp>foo</samp>"},
		{Script("alert('foo')"), "<script>alert('foo')</script>"},
		{ScriptSrc("/main.js"), "<script src=\"/main.js\" />"},
		{Section(Text("foo")), "<section>foo</section>"},
		{Select(Text("foo")), "<select>foo</select>"},
		{Small(Text("foo")), "<small>foo</small>"},
		{Source(Text("foo")), "<source>foo</source>"},
		{Span(Text("foo")), "<span>foo</span>"},
		{Strong(Text("foo")), "<strong>foo</strong>"},
		{Style("body{background:red}"), "<style>body{background:red}</style>"},
		{Sub(Text("foo")), "<sub>foo</sub>"},
		{Summary(Text("foo")), "<summary>foo</summary>"},
		{Sup(Text("foo")), "<sup>foo</sup>"},
		{Svg(Text("foo")), "<svg>foo</svg>"},
		{Table(Text("foo")), "<table>foo</table>"},
		{Tbody(Text("foo")), "<tbody>foo</tbody>"},
		{Td(Text("foo")), "<td>foo</td>"},
		{Template(Text("foo")), "<template>foo</template>"},
		{Textarea(Text("foo")), "<textarea>foo</textarea>"},
		{Tfoot(Text("foo")), "<tfoot>foo</tfoot>"},
		{Th(Text("foo")), "<th>foo</th>"},
		{Thead(Text("foo")), "<thead>foo</thead>"},
		{Time(Text("foo")), "<time>foo</time>"},
		{Title("hello world"), "<title>hello world</title>"},
		{Tr(Text("foo")), "<tr>foo</tr>"},
		{Track(Text("foo")), "<track>foo</track>"},
		{U(Text("foo")), "<u>foo</u>"},
		{Ul(Text("foo")), "<ul>foo</ul>"},
		{Var(Text("foo")), "<var>foo</var>"},
		{Video(Text("foo")), "<video>foo</video>"},
		{Wbr(Text("foo")), "<wbr>foo</wbr>"},
	}

	for _, testCase := range testCases {
		result := RenderOpts(testCase.node, true, nil)
		assertEqual(t, testCase.expected, result)
	}
}
