package main

import (
	"fmt"

	s "github.com/oetherington/smetana"
)

func main() {
	node := s.Html(
		s.Attrs{},
		s.Children{
			s.Title("My HTML Document"),
			s.Link("stylesheet", "/styles/index.css"),
		},
		s.Children{
			s.Div(s.Attrs{"class": "hello"}, s.Children{
				s.Span(s.Attrs{}, s.Children{
					s.Text("Hello world"),
				}),
			}),
		},
	)
	fmt.Println(s.Render(node))
}
