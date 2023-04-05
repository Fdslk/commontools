package cmd

import (
	"testing"
)

func TestPwdGeneratorGenerateValidPasswordWhenUseDefaultCondtion(t *testing.T) {
	testpwd := "36la9x5:WS>r?>8_"
	valid := validatePassword(testpwd)
	if !valid {
		t.Errorf("Expected true but got %t", valid)
	}
}

func TestPwdGeneratorGenerateValidPasswordWhenUseDefaultCondtion2(t *testing.T) {
	testpwd := "36la9x5:WS>r?>8_"
	containDigit = true
	containLowerCase = true
	containUpperCase = true
	containSpecialCase = true
	valid := validatePassword(testpwd)
	if !valid {
		t.Errorf("Expected true but got %t", valid)
	}
}

func TestPwdGeneratorGenerateInValidPasswordWhenShouldIncludeDigit(t *testing.T) {
	testpwd := "LRRcG<b]k_#/Y"
	containDigit = true
	valid := validatePassword(testpwd)
	if valid {
		t.Errorf("Expected false but got %t", valid)
	}
}

func TestPwdGeneratorGenerateInValidPasswordWhenShouldIncludeUpperCase(t *testing.T) {
	testpwd := "fdeafdsafdsaf"
	containUpperCase = true
	valid := validatePassword(testpwd)
	if valid {
		t.Errorf("Expected false but got %t", valid)
	}
}

func TestPwdGeneratorGenerateInValidPasswordWhenShouldIncludeLowerCase(t *testing.T) {
	testpwd := "QQQQQQQQQQQQ"
	containLowerCase = true
	valid := validatePassword(testpwd)
	if valid {
		t.Errorf("Expected false but got %t", valid)
	}
}

func TestPwdGeneratorGenerateInValidPasswordWhenShouldIncludeSpecialCase(t *testing.T) {
	testpwd := "seafhjkeh12312JKAJHJ"
	containSpecialCase = true
	valid := validatePassword(testpwd)
	if valid {
		t.Errorf("Expected false but got %t", valid)
	}
}
