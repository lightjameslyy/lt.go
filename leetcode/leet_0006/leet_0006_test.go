package leet_0006

import (
	"testing"
)

// 表格驱动测试
var tests = []struct {
	s       string
	numRows int
	res     string
}{
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	{"PAYPALISHIRING", 5, "PHASIYIRPLIGAN"},
}

func TestConvert(t *testing.T) {
	for _, tt := range tests {
		res := Convert(tt.s, tt.numRows)
		if tt.res != res {
			t.Errorf("test: %v, expected: %v, got %v",
				tt, tt.res, res)
		}
	}
}
