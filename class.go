package smetana

import "strings"

type Classes map[string]bool

func Class(args ...any) string {
	classes := []string{}
	for _, arg := range args {
		switch item := arg.(type) {
		case string:
			classes = append(classes, item)
		case Classes:
			for key, value := range item {
				if value {
					classes = append(classes, key)
				}
			}
		default:
			break
		}
	}
	return strings.Join(classes, " ")
}
