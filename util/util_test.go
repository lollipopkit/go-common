package util_test

import (
	"testing"

	"github.com/lollipopkit/gommon/util"
)

func TestContains(t *testing.T) {
	l := []string{"1"}
	s1 := "1"
	s2 := "2"
	if !util.Contains(l, s1) {
		t.Fatal("Contains s1 failed")
	}
	if util.Contains(l, s2) {
		t.Fatal("Contains s2 failed")
	}
}

func TestClear(t *testing.T) {
	l := []string{"1"}
	util.Clear(&l)
	if len(l) != 0 {
		t.Fatal("Clear failed")
	}
}