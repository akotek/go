package tests

import (
	"go/main/training"
	"testing"
)

func TestAnagram(t *testing.T) {

	if training.Anagram("listen", "silent") != true{
		t.Error("expected true")
	}
	if training.Anagram("triangle", "integral") != true{
		t.Error("expected true")
	}
	if training.Anagram("rdbms", "rdbm") != false{
		t.Error("expected false")
	}
}