package main

import (
	"fmt"
	s "github.com/oetherington/smetana"
)

func main() {
	styles := s.NewStyleSheet()
	font := styles.AddFont("OpenSans", "OpenSans.woff2")
	container := styles.AddAnonClass(s.CssProps{
		"font-family": font,
		"background":  s.Rgb(255, 0, 255),
		"padding":     s.EM(2),
	})

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
				s.Span(s.Text("Hello world")),
			),
			s.Div(s.Text("foobar")),
		),
	)

	fmt.Println("\nGenerated CSS:")
	fmt.Println(s.RenderCss(styles, s.Palette{}))
	fmt.Println("\nGenerated HTML:")
	fmt.Println(s.RenderHtml(node))
}
