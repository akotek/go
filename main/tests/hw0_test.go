package tests

import (
	"go/main/training"
	"testing"
)

func TestPrime(t *testing.T){
	if training.IsPrime(3) != true && training.IsPrime(89) != true{
		t.Error("expected true")
	}
	if training.IsPrime(4) != false && training.IsPrime(60) != false{
		t.Error("expected false")
	}
}


func TestPal(t *testing.T){
	// even test
	if training.IsPalindrome("abba") != true{
		t.Error("expected true")
	}
	// odd test
	if training.IsPalindrome("radar") != true{
		t.Error("expected true")
	}
	if training.IsPalindrome("a") != true{
		t.Error("expected true")
	}
	if training.IsPalindrome("radzr") != false{
		t.Error("expected false")
	}
	if training.IsPalindrome("abbbbbbca") != false{
		t.Error("expected false")
	}
}
