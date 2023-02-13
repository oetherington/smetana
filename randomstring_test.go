package smetana

import "testing"

func TestRandomStringIsRandom(t *testing.T) {
	s1 := randomString(8)
	s2 := randomString(8)
	assertNotEqual(t, s1, s2)
}

func TestRandomStringHasCorrectLength(t *testing.T) {
	s1 := randomString(8)
	s2 := randomString(47)
	assertEqual(t, 8, len(s1))
	assertEqual(t, 47, len(s2))
}
