package smetana

import "testing"

func TestMergeMaps(t *testing.T) {
	m1 := map[string]int{"foo": 1, "bar": 2}
	m2 := map[string]int{"baz": 3}
	MergeMaps(m1, m2)
	assertEqual(t, map[string]int{"foo": 1, "bar": 2, "baz": 3}, m1)
	assertEqual(t, map[string]int{"baz": 3}, m2)
}

func TestMergeMapsOverwritesDuplicates(t *testing.T) {
	m1 := map[string]int{"foo": 1, "bar": 2, "baz": 4}
	m2 := map[string]int{"baz": 3}
	MergeMaps(m1, m2)
	assertEqual(t, map[string]int{"foo": 1, "bar": 2, "baz": 3}, m1)
	assertEqual(t, map[string]int{"baz": 3}, m2)
}

func TestIdHelper(t *testing.T) {
	result := Id("foo")
	assertEqual(t, Attr{Key: "id", Value: "foo"}, result)
}

func TestMin(t *testing.T) {
	assertEqual(t, 2, min(2, 3))
	assertEqual(t, 2, min(3, 2))
}

func TestMax(t *testing.T) {
	assertEqual(t, 3, max(2, 3))
	assertEqual(t, 3, max(3, 2))
}

func TestClamp(t *testing.T) {
	assertEqual(t, 15, clamp(15, 10, 20))
	assertEqual(t, 10, clamp(5, 10, 20))
	assertEqual(t, 10, clamp(10, 10, 20))
	assertEqual(t, 20, clamp(25, 10, 20))
	assertEqual(t, 20, clamp(20, 10, 20))
}

func TestXform(t *testing.T) {
	values := []int32{1, -2, 3}
	result := Xform(values, func(a int32) uint32 {
		if a < 0 {
			a = -a
		}
		return uint32(a)
	})
	assertEqual(t, []uint32{1, 2, 3}, result)
}
