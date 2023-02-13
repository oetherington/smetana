package main

import (
	"fmt"

	s "github.com/oetherington/smetana"
)

func main() {
	node := s.Html(
		s.Head(
			s.Charset(""),
			s.Title("My HTML Document"),
			s.Link("stylesheet", "/styles/index.css"),
			s.Keywords("smetana,template,rendering,golang"),
			s.Description("Smetana templates for Golang"),
			s.Author("Ollie Etherington"),
			s.Viewport(""),
		),
		s.Body(
			s.Div(
				s.Attrs{"class": s.ClassNames("foo", "bar")},
				s.Span(s.Text("Hello world")),
			),
			s.Div(s.Text("foobar")),
		),
	)
	fmt.Println(s.Render(node))
}
