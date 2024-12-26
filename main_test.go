package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	input string
	valid bool
}

func TestSolve(t *testing.T) {
	inputs := []string{
		"031702864607081030008900100302640000185070603096000050720103508009024017000590400",
		"100097000000000576427806900000900035000060000509040002004630800008010700902085000",
		"600030000000009060500000109000702003200080600430000801305204000000600070040007000",
		"013609700000000040090000000000400601080100200000007050002001087705000004000040000",
		"000000000830010400000070800904030100005940060000005020060000000400800000209000710",
		"000000000830010400000070800904030100005940060000005020060000000400800000209000710",
		"400009500290750000503000004000030020600007080040200005000000000376400000000080607",
	}

	for i, input := range inputs {
		grid := load(input, 9)
		solve(&grid)
		result := validate(&grid)
		if !result {
			t.Errorf("Expected %t, got %t. index: %d\n", true, result, i)
			render(&grid)
			fmt.Println("")
		}
	}
}

func TestValidate(t *testing.T) {
	inputs := []testCase{
		{
			input: "031702864607081030008900100302640000185070603096000050720103508009024017000590400",
			valid: false,
		},
		{
			input: "487369512291754368563128794719835426625947183843216975158673249376492851492581637",
			valid: false,
		},
	}

	for _, tc := range inputs {
		grid := load(tc.input, 9)
		result := validate(&grid)
		if result != tc.valid {
			t.Errorf("Expected %t, got %t", tc.valid, result)
		}
	}
}
