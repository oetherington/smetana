package smetana

import "testing"

func TestCanDetectVoidTags(t *testing.T) {
	assertEqual(t, true, isVoidTag("area"))
	assertEqual(t, true, isVoidTag("col"))
	assertEqual(t, true, isVoidTag("wbr"))
	assertEqual(t, false, isVoidTag("div"))
	assertEqual(t, false, isVoidTag("span"))
}
