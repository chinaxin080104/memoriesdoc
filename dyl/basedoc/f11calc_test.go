package main

import "testing"

//单独运行，会报错。

func TestAddfuc119(*testing.T) {
	if ans := Addfuc119(1, 2); ans != 3 {
		// 如果不符合预期，那就是测试不通过
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}
}
