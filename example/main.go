package main

import (
	"fmt"
	s "github.com/oetherington/smetana"
)

func main() {
	smetana := s.NewSmetanaWithPalettes(s.Palettes{
		"light": {
			"bg": s.Hex("#eee"),
			"fg": s.Hex("#222"),
		},
		"dark": {
			"bg": s.Hex("#363636"),
			"fg": s.Hex("#ddd"),
		},
	})

	font := smetana.Styles.AddFont("OpenSans", "/OpenSans.woff2")
	container := smetana.Styles.AddAnonClass(s.CssProps{
		{Key: "font-family", Value: font},
		{Key: "padding", Value: s.EM(2)},
		{Key: "background", Value: s.PaletteValue("bg")},
		{Key: "color", Value: s.PaletteValue("fg")},
	})

	css := smetana.RenderStyles()

	node := s.Html(
		s.Head(
			s.Charset(""),
			s.Title("My HTML Document"),
			s.LinkHref("stylesheet", "/styles/index.css"),
			s.Keywords("smetana,template,rendering,golang"),
			s.Description("Smetana templates for Golang"),
			s.Author("Ollie Etherington"),
			s.Viewport(""),
		),
		s.Body(
			container,
			s.Div(
				s.Attrs{"aria-label": "hello world"},
				s.ClassNames("foo", "bar"),
				s.Span("Hello world"),
			),
			s.Div("foobar"),
		),
	)

	html := s.RenderHtml(node)

	fmt.Println("\nGenerated CSS (light mode):")
	fmt.Println(css["light"])

	fmt.Println("\nGenerated CSS (dark mode):")
	fmt.Println(css["dark"])

	fmt.Println("\nGenerated HTML:")
	fmt.Println(html)
}
