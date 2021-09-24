package main

import "testing"

func TestHWString(t *testing.T) {
	if HWString() != "hello world" {
		t.Errorf("bukan %q, seharusnya hello world", HWString())
	}
}
