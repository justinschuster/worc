package utils

import (
	"testing"
)

func TestParseTocFile(t *testing.T) {
	bad_path := "/home/justin/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft/_retail_/Interface/Addons/WeakAuras/"
	pattern := `^## Version: (.+)`

	result := ParseTocFile(bad_path, pattern)
	if result == "" {
		t.Fatalf("Bad path ParseTocFile")
	}
}
