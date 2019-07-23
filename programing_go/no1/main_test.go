package main

import (
	"flag"
	"testing"
)

func Test_programを起動したコマンド名であるarg0を表示する(t *testing.T) {
	// 前準備
	flag.CommandLine.Set("args0", "test1")

	// 実行
	arr := practice1()

	// 検証
	if arr == nil && len(arr) < 1 {
		t.Error("Don't get array")
	}

	t.Log(arr)
}
