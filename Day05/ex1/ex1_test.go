package main

import "testing"

func TestCheck(t *testing.T) {
	bp := BoardingPass{rows: Half{0, 127}, cols: Half{0, 7}}

	or := bp.rows
	oc := bp.cols
	r, c, s := bp.Check("FBFBBFFRLR")
	if r != 44 || c != 5 || s != 357 {
		t.Errorf("Test faild. Got %d %d %d, epxecte 44 4 356",
			r, c, s)
	}

	bp.rows = or
	bp.cols = oc
	r, c, s = bp.Check("BFFFBBFRRR")
	if r != 70 || c != 7 || s != 567 {
		t.Errorf("Test faild. Got %d %d %d, epxecte 44 4 356",
			r, c, s)
	}

	bp.rows = or
	bp.cols = oc
	r, c, s = bp.Check("FFFBBBFRRR")
	if r != 14 || c != 7 || s != 119 {
		t.Errorf("Test faild. Got %d %d %d, epxecte 44 4 356",
			r, c, s)
	}

	bp.rows = or
	bp.cols = oc
	r, c, s = bp.Check("BBFFBBFRLL")
	if r != 102 || c != 4 || s != 820 {
		t.Errorf("Test faild. Got %d %d %d, epxecte 44 4 356",
			r, c, s)
	}
}
