package fromstr

import "testing"

type testCaseInt64 struct {
	Str string
	Def int64
	Res int64
}

type testCaseFloat64 struct {
	Str string
	Def float64
	Res float64
}

func TestInt64(t *testing.T) {

	testCases := []testCaseInt64{
		testCaseInt64{"1", 0, 1},
		testCaseInt64{"-1", 0, -1},
		testCaseInt64{"1.23", 0, 0},
		testCaseInt64{"0", 1, 0},
		testCaseInt64{"a", 1, 1},
		testCaseInt64{"a", 0, 0},
		testCaseInt64{"1a", 2, 2},
		testCaseInt64{" ", 2, 2},
	}

	for _, x := range testCases {
		if Int64(x.Str, x.Def) != x.Res {
			t.Errorf(
				"Int64 was given str %s, def %d and expected %d, got %d",
				x.Str,
				x.Def,
				x.Res,
				Int64(x.Str, x.Def))
		}
	}
}

func TestFloat64(t *testing.T) {

	testCases := []testCaseFloat64{
		testCaseFloat64{"1", 0.0, 1},
		testCaseFloat64{"1.23", 0.0, 1.23},
		testCaseFloat64{"-1.23", 0.0, -1.23},
		testCaseFloat64{"0", 1, 0},
		testCaseFloat64{"a", 1, 1},
		testCaseFloat64{"a", 0, 0},
		testCaseFloat64{"1a", 2, 2},
		testCaseFloat64{" ", 2, 2},
	}

	for _, x := range testCases {
		if Float64(x.Str, x.Def) != x.Res {
			t.Errorf(
				"Int64 was given str %s, def %f and expected %f, got %f",
				x.Str,
				x.Def,
				x.Res,
				Float64(x.Str, x.Def))
		}
	}
}
