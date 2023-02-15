package smetana

import "fmt"

// A map from CSS property names to their values. For instance,
//
//	{"background": "red"}.
//
// For properties that take size values it is recommended to use a
// unit helper rather than setting the value with a string. For
// example, instead of
//
//	{"padding": "4px"}
//
// you can use
//
//	{"padding": PX(4)}
type CssProps map[string]any

// The name of a CSS class.
type ClassName string

// [StyleSheet] aggregates the CSS styles for a page and compiles them
// from the in-code representation into a CSS string for the browser.
type StyleSheet struct {
	Classes map[ClassName]CssProps
}

// Create a new empty [StyleSheet].
func NewStyleSheet() StyleSheet {
	return StyleSheet{map[ClassName]CssProps{}}
}

// Add a new class to a [StyleSheet].
func (styles *StyleSheet) AddClass(props CssProps) ClassName {
	name := ClassName(RandomString(8))
	styles.Classes[name] = props
	return name
}

// Compile a [StyleSheet] into a CSS String.
func (styles StyleSheet) ToCss(builder *Builder) {
	for name, props := range styles.Classes {
		builder.Buf.WriteByte('.')
		builder.Buf.WriteString(string(name))
		builder.Buf.WriteByte('{')
		writeProps(builder, props)
		builder.Buf.WriteByte('}')
	}
}

func writeProps(builder *Builder, props CssProps) {
	for key, value := range props {
		builder.Buf.WriteString(key)
		builder.Buf.WriteByte(':')

		switch item := value.(type) {
		case string:
			builder.Buf.WriteString(item)
		case fmt.Stringer:
			builder.Buf.WriteString(item.String())
		case Color:
			builder.Buf.WriteString(item.ToCssColor())
		case int:
			builder.Buf.WriteString(fmt.Sprintf("%dpx", item))
		default:
			err := fmt.Errorf("Invalid CSS value: %v", item)
			builder.Logger.Println(err)
		}

		builder.Buf.WriteByte(';')
	}
}
