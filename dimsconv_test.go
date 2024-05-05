package dimsconv

import (
	"github.com/unixlinuxgeek/floatgen"
	"testing"
)

func TestInchToCm(t *testing.T) {
	r := Inch(floatgen.GenRan(1, 9))
	f := Cm(r * 2.54)

	if InchToCm(r) == f {
		t.Logf("Pass: InchToCm(%f) == %f\n", r, f)
	} else {
		t.Errorf("Error: InchToCm(%f) != %f\n", r, f)
	}
}

func TestCmToInch(t *testing.T) {
	r := Cm(floatgen.GenRan(1, 9))
	f := Inch(r / 2.54)

	if CmToInch(r) == f {
		t.Logf("Pass: CmToInch(%f) == %f\n", r, f)
	} else {
		t.Errorf("Error: CmToInch(%f) != %f\n", r, f)
	}
}
