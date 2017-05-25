package utils

import (
	"testing"
)

func TestIsDigitCharacter(t *testing.T) {
	if !IsDigitCharacter("2") {
		t.Errorf("IsDigitCharacter(\"2\") : Incorrect result. Expected (true), Found (false).")
	}

	if !IsDigitCharacter("0") {
		t.Errorf("IsDigitCharacter(\"0\") : Incorrect result. Expected (true), Found (false).")
	}

	if IsDigitCharacter("r") {
		t.Errorf("IsDigitCharacter(\"r\") : Incorrect result. Expected (false), Found (true).")
	}

	if IsDigitCharacter("E") {
		t.Errorf("IsDigitCharacter(\"E\") : Incorrect result. Expected (false), Expected (true).")
	}

	if IsDigitCharacter("11") {
		t.Errorf("IsDigitCharacter(\"11\") : Incorrect result. Expected (false), Expected (true).")
	}
}

func TestIsAlphaCharacter(t *testing.T) {
	if !IsAlphaCharacter("a") {
		t.Errorf("IsAlphaCharacter(\"a\") : Incorrect result. Expected (true), Found (false).")
	}

	if !IsAlphaCharacter("A") {
		t.Errorf("IsAlphaCharacter(\"A\") : Incorrect result. Expected (true), Found (false).")
	}

	if !IsAlphaCharacter("_") {
		t.Errorf("IsAlphaCharacter(\"_\") : Incorrect result. Expected (true), Found (false).")
	}

	if IsAlphaCharacter("5") {
		t.Errorf("IsAlphaCharacter(\"5\") : Incorrect result. Expected (false), Found (true).")
	}

	if IsAlphaCharacter("aB") {
		t.Errorf("IsAlphaCharacter(\"aB\") : Incorrect result. Expected (false), Found (true).")
	}
}

func TestIsAlphaNumericCharacter(t *testing.T) {
	if !IsAlphaNumericCharacter("a") {
		t.Errorf("IsAlphaNumericCharacter(\"a\") : Incorrect result. Expected (true), Found (false).")
	}

	if !IsAlphaNumericCharacter("A") {
		t.Errorf("IsAlphaNumericCharacter(\"A\") : Incorrect result. Expected (true), Found (false).")
	}

	if !IsAlphaNumericCharacter("_") {
		t.Errorf("IsAlphaNumericCharacter(\"_\") : Incorrect result. Expected (true), Found (false).")
	}

	if !IsAlphaNumericCharacter("5") {
		t.Errorf("IsAlphaNumericCharacter(\"5\") : Incorrect result. Expected (true), Found (false).")
	}

	if IsAlphaNumericCharacter("a3") {
		t.Errorf("IsAlphaNumericCharacter(\"a3\") : Incorrect result. Expected (false), Found (true).")
	}
}
