package smetana

import "strings"

// A map of conditional classes to be passed to ClassNames. The keys
// are class names and the values are booleans indicating whether or
// not to include that class name. For instance,
//
//	{"foo": true, "bar" false}
//
// will evaluate to "foo".
type Classes map[ClassName]bool

// A utility function for concatenating multiple class names into a
// single string suitable for embedding in HTML. Arguments may be of
// several different types:
//   - string
//   - [ClassName]
//   - [Classes]
//
// Arguments of other types are ignored.
func ClassNames(args ...any) string {
	classes := []string{}
	for _, arg := range args {
		switch item := arg.(type) {
		case string:
			if len(item) > 0 {
				classes = append(classes, item)
			}
		case ClassName:
			if len(item) > 0 {
				classes = append(classes, string(item))
			}
		case Classes:
			for key, value := range item {
				if value {
					classes = append(classes, string(key))
				}
			}
		default:
			break
		}
	}
	return strings.Join(classes, " ")
}
