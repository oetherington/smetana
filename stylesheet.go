package smetana

import (
	"fmt"
	"strings"
)

type CssProps map[string]any

type ClassName string

type StyleSheet struct {
	Classes map[ClassName]CssProps
}

func NewStyleSheet() StyleSheet {
	return StyleSheet{map[ClassName]CssProps{}}
}

func (styles *StyleSheet) AddClass(props CssProps) ClassName {
	name := ClassName(randomString(8))
	styles.Classes[name] = props
	return name
}

func (styles StyleSheet) ToCss() string {
	var sb strings.Builder

	for name, props := range styles.Classes {
		sb.WriteByte('.')
		sb.WriteString(string(name))
		sb.WriteByte('{')
		writeProps(&sb, props)
		sb.WriteByte('}')
	}

	return sb.String()
}

func writeProps(sb *strings.Builder, props CssProps) {
	for key, value := range props {
		sb.WriteString(key)
		sb.WriteByte(':')

		switch item := value.(type) {
		case string:
			sb.WriteString(item)
		case fmt.Stringer:
			sb.WriteString(item.String())
		case int:
			sb.WriteString(fmt.Sprintf("%dpx", item))
		default:
			break // TODO: How should we handle this error?
		}

		sb.WriteByte(';')
	}
}
