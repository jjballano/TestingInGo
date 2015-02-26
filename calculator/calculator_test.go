package calculator

import "testing"

func TestSumsTwoNumbers(t *testing.T) {
	sumValue := Sum(3,5)
	if sumValue != 8 {
		t.Errorf("The result of %d + %d must be %d",3,5,8) 
	}
}