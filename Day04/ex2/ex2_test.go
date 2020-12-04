package main

import "testing"

func TestValidPassportsInvalid(t *testing.T) {
	result := 0
	p := PassportScanner{}
	p.readPassports("../input_testinvalid.dat")

	if p.nValid != result {
		t.Errorf("Test failed. Got %d and expected %d.",
			p.nValid, result)
	}
}

func TestValidPassportsValid(t *testing.T) {
	result := 4
	p := PassportScanner{}
	p.readPassports("../input_testvalid.dat")

	if p.nValid != result {
		t.Errorf("Test failed. Got %d and expected %d.",
			p.nValid, result)
	}
}

func TestValidField(t *testing.T) {
	p := PassportScanner{}
	// byr valid:   2002
	valid := p.validateField("byr", "2002")
	if !valid {
		t.Errorf("Test faild. byr 2002 expected valid got %t",
			valid)
	}
	// byr invalid: 2003
	valid = p.validateField("byr", "2003")
	if valid {
		t.Errorf("Test faild. byr 2003 expected invalid got %t",
			valid)
	}
	// eyr invalid: 1967
	valid = p.validateField("eyr", "1967")
	if valid {
		t.Errorf("Test faild. eyr 1967 expected invalid got %t",
			valid)
	}
	//
	// hgt valid:   60in
	valid = p.validateField("hgt", "60in")
	if !valid {
		t.Errorf("Test faild. hgt 60in expected valid got %t",
			valid)
	}
	// hgt valid:   190cm
	valid = p.validateField("hgt", "190cm")
	if !valid {
		t.Errorf("Test faild. hgt 190cm expected valid got %t",
			valid)
	}
	// hgt invalid: 190in
	valid = p.validateField("hgt", "190in")
	if valid {
		t.Errorf("Test faild. hgt 190in expected invalid got %t",
			valid)
	}
	// hgt invalid: 190
	valid = p.validateField("hgt", "190")
	if valid {
		t.Errorf("Test faild. hgt 190 expected invalid got %t",
			valid)
	}
	//
	// hcl valid:   #123abc
	valid = p.validateField("ecl", "brn")
	if !valid {
		t.Errorf("Test faild. ecl brn expected valid got %t",
			valid)
	}
	// hcl invalid: #123abz
	valid = p.validateField("hcl", "123abc")
	if valid {
		t.Errorf("Test faild. hcl 123abc expected invalid got %t",
			valid)
	}

	// hcl invalid: 123abc
	valid = p.validateField("hcl", "123abc")
	if valid {
		t.Errorf("Test faild. hcl 123abc expected invalid got %t",
			valid)
	}
	//
	// ecl valid:   brn
	valid = p.validateField("ecl", "brn")
	if !valid {
		t.Errorf("Test faild. ecl brn expected valid got %t",
			valid)
	}
	// ecl invalid: wat
	valid = p.validateField("ecl", "wat")
	if valid {
		t.Errorf("Test faild. ecl wat expected invalid got %t",
			valid)
	}
	//
	// pid valid:   000000001
	valid = p.validateField("pid", "000000001")
	if !valid {
		t.Errorf("Test faild. pid 000000001 expected valid got %t",
			valid)
	}
	// pid invalid: 0123456789
	valid = p.validateField("pid", "0123456789")
	if valid {
		t.Errorf("Test faild. pid 0123456789 expected invalid got %t",
			valid)
	}

}
