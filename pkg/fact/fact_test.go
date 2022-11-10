package fact

import "testing"

func TestFact(t *testing.T) {
	testTable := []struct {
		x    int
		must int
	}{
		{x: 3, must: 6},
		{x: 4, must: 24},
		{x: 6, must: 720},
	}

	for _, test := range testTable {
		res := Fact(test.x)

		if res != test.must {
			t.Error("wrong fact result. X, Res, Mest:", test.x, res, test.must)
		}
	}
}
