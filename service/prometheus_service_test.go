package service

import (
	"testing"
)

func TestReplaceInvalidChar(t *testing.T) {
	expected := "aaa_bbb"

	testString := "aaa.bbb"
	got := ReplaceInvalidChar(testString)
	if got != expected {
		t.Errorf("%s => %s; want %s", testString, got, expected)
	}

	testString = "aaa-bbb"
	got = ReplaceInvalidChar(testString)
	if got != expected {
		t.Errorf("%s => %s; want %s", testString, got, expected)
	}
}
