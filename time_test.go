package main

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
	"time"
)

func TestTime01(t *testing.T) {
	var f float64 = 1670921527.8599809
	timestamp := cast.ToInt64(f * 1000)
	fmt.Println(timestamp)
	fmt.Println(time.UnixMilli(timestamp))
	// 2022-12-13 16:52:07 +0800 CST
	// 2022-12-13 16:52:07.859 +0800 CST

}
