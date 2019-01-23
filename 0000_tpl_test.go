package test

import (
	"testing"
)

// en name
// zh name
// 难度 简单
//
// desc

func TestL000(t *testing.T) {
	if tpl("alex", "aaleex") != true {
		t.Fail()
	}
}

// 思路
func tpl(name string, typed string) bool {
	return true
}
