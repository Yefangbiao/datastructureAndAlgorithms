package bagQuestion

import "testing"

func TestBag0_1(t *testing.T) {
	tests := []struct {
		n, c  int
		goods []Goods
	}{
		{5, 10, []Goods{
			Goods{Weight: 2, Value: 6}, Goods{Weight: 2, Value: 3}, Goods{Weight: 6, Value: 5},
			Goods{Weight: 5, Value: 4}, Goods{Weight: 4, Value: 6},
		}},
	}

	for _, test := range tests {
		got := MaxValue(test.n, test.c, test.goods)
		want := 15
		if got != want {
			t.Errorf("error ,got: %v, want: %v", got, want)
		}
		got = MaxValue2(test.n, test.c, test.goods)
		if got != want {
			t.Errorf("error ,got: %v, want: %v", got, want)
		}
	}

}
