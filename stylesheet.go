package smetana

import (
	"fmt"
	"strings"
)

/*
 * A map from CSS property names to their values. For instance,
 * {"background": "red"}. For properties that take size values it
 * is recommended to use a unit helper rather than setting the value
 * with a string. For example, instead of {"padding": "4px"}, you can
 * use {"padding": PX(4)}.
 */
type CssProps map[string]any

// The name of a CSS class
type ClassName string

/*
 * StyleSheet aggregates the CSS styles for a page and compiles them
 * from the in-code representation into a CSS string for the browser.
 */
type StyleSheet struct {
	Classes map[ClassName]CssProps
}

// Create a new empty StyleSheet
func NewStyleSheet() StyleSheet {
	return StyleSheet{map[ClassName]CssProps{}}
}

// Add a new class to a StyleSheet
func (styles *StyleSheet) AddClass(props CssProps) ClassName {
	name := ClassName(randomString(8))
	styles.Classes[name] = props
	return name
}

// Compile a StyleSheet into a CSS String
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
