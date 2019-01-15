package test

import (
	"testing"
	"strconv"
	"bytes"
)

// Fizz Buzz
//写一个程序，输出从 1 到 n 数字的字符串表示。
//
//1. 如果 n 是3的倍数，输出“Fizz”；
//
//2. 如果 n 是5的倍数，输出“Buzz”；
//
//3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

func TestL412(t *testing.T) {
	buzz := fizzBuzz(1)
	if !compareStringSlice(buzz, []string{"1"}) {
		t.Fail()
	}
}

// 2019/1/15
// 这个题太简单了, 最近今天忙, 先刷点简单的题
func fizzBuzz(n int) []string {
	ss := make([]string, n)
	bf := bytes.Buffer{}
	for i := 1; i < n+1; i++ {
		if i%3 == 0 {
			bf.WriteString("Fizz")
		}
		if i%5 == 0 {
			bf.WriteString("Buzz")
		}
		if bf.Len() == 0 {
			bf.WriteString(strconv.Itoa(i))
		}
		ss[i-1] = bf.String()
		bf.Reset()
	}

	return ss
}
