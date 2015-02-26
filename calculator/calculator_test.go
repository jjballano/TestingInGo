package calculator

import "testing"

func TestSumsTwoNumbers(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{2,3,5},
		{0,0,0},
		{0,1,1},
		{3,0,3},		
		{4,-2,2},
		{-1,-3,-4},
		{-1,4,3},		
	}
	for _,numbers := range cases {
		sum := Sum(numbers.a, numbers.b)
		if (sum != numbers.expected){
			t.Errorf("The result of %d + %d must be %d, but was %d",numbers.a,numbers.b,numbers.expected, sum) 	
		}
	}

}