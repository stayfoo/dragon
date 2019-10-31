/*
@Time : 2019-10-31 12:57
@Author : mengyueping
@File : calculator_test
@Software: GoLand
*/
package calculator

import "testing"

func TestEvaluate(t *testing.T) {
	tests := []struct{
		input string
		expected int
	}{
		{"1+1;",2},
		{"2*3;",6},
		{"4/2;",2},
		{"1+2+3+4;",10},
		{"2*3+4;",10},
		{"1+2*3;",7},
		{"1+2*3+4;",11},
		{"1+2*3+4/2;",9},
		{"1+2*3+4/2-1;",8},
	}

	for _, value := range tests {
		cal := Calculator{}
		cal.Code = value.input
		cal.Evaluate()
		if cal.Result != value.expected {
			t.Errorf("Wrong answer, got=%d, want=%d", cal.Result, value.expected)
		}
	}
}
