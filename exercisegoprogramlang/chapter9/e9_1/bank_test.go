package bank

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	var tests = []struct {
		deposit  int
		withdraw int
		want     bool
	}{
		{
			100, 101, false,
		},
		{
			99, 100, true,
		},
	}
	for _, test := range tests {
		Deposit(test.deposit)
		if got := Withdraw(test.withdraw); got != test.want {
			t.Errorf("error! got = %v, want = %v\n", got, test.want)
		}
	}
}
