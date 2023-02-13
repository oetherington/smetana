package main

import (
	"fmt"

	s "github.com/oetherington/smetana"
)

func main() {
	node := s.Html(
		s.Attrs{},
		s.Children{
			s.Charset(""),
			s.Title("My HTML Document"),
			s.Link("stylesheet", "/styles/index.css"),
			s.Keywords("smetana,template,rendering,golang"),
			s.Description("Smetana templates for Golang"),
			s.Author("Ollie Etherington"),
			s.Viewport(""),
		},
		s.Children{
			s.Div(s.Attrs{"class": s.Class("foo", "bar")}, s.Children{
				s.Span(s.Attrs{}, s.Children{
					s.Text("Hello world"),
				}),
			}),
		},
	)
	fmt.Println(s.Render(node))
}
