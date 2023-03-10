package smetana

import "testing"

func TestClassNamesConcatsStringsAndClassNames(t *testing.T) {
	result := ClassNames("foo", "bar", ClassName("baz"))
	assertEqual(t, "foo bar baz", result)
}

func TestClassNamesIgnoresEmptyStrings(t *testing.T) {
	result := ClassNames("foo", "", "bar", "")
	assertEqual(t, "foo bar", result)
}

func TestClassNamesCanbeEmpty(t *testing.T) {
	result := ClassNames()
	assertEqual(t, "", result)
}

func TestClassesNamesCanBeConditional(t *testing.T) {
	one := 1
	result := ClassNames(Classes{
		"a": true,
		"b": false,
		"c": 1 == 2,
		"d": 1 == one,
	})
	assertOneOf(t, []ClassName{"a d", "d a"}, result)
}

func TestCanMixConditionalAndUnconditionalClasses(t *testing.T) {
	result := ClassNames("foo", Classes{
		"a": true,
		"b": false,
	}, "bar", Classes{
		"e": false,
		"f": true,
	})
	assertEqual(t, "foo a bar f", result)
}

func TestClassNamesIgnoresInvalidArgumentTypes(t *testing.T) {
	result := ClassNames("foo", 3)
	assertEqual(t, "foo", result)
}
