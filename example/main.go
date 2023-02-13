package main

import (
	"fmt"
	s "github.com/oetherington/smetana"
)

func main() {
	styles := s.NewStyleSheet()
	container := styles.AddClass(s.CssProps{
		"background": "red",
		"padding":    s.EM(2),
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
				s.Attrs{"class": s.ClassNames("foo", "bar")},
				s.Span(s.Text("Hello world")),
			),
			s.Div(s.Text("foobar")),
		),
	)

	fmt.Println(styles.ToCss())
	fmt.Println(s.Render(node))
}
