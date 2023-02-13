package smetana

import "testing"

func TestClassConcatsStrings(t *testing.T) {
	result := Class("foo", "bar", "baz")
	assertEqual(t, "foo bar baz", result)
}

func TestClassCanbeEmpty(t *testing.T) {
	result := Class()
	assertEqual(t, "", result)
}

func TestClassesCanBeConditional(t *testing.T) {
	result := Class(Classes{
		"a": true,
		"b": false,
		"c": 1 == 2,
		"d": 1 == 1,
	})
	assertOneOf(t, []string{"a d", "d a"}, result)
}

func TestCanMixConditionalAndUnconditionalClasses(t *testing.T) {
	result := Class("foo", Classes{
		"a": true,
		"b": false,
	}, "bar", Classes{
		"e": false,
		"f": true,
	})
	assertEqual(t, "foo a bar f", result)
}
