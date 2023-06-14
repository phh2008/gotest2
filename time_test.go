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

func TestRFC3339(t *testing.T) {
	ts, _ := time.Parse(time.RFC3339, "2023-01-09T09:58:45.368895827Z")
	local, _ := time.LoadLocation("Asia/Shanghai")
	date := ts.In(local).Format("2006-01-02 15:04:05")
	fmt.Println(date)
}

// UTC 转北京时间
func TestUTC2CST(t *testing.T) {
	val := "2023-01-09 09:54:45.3580496 +0000 UTC"
	layout := "2006-01-02 15:04:05.0000000 +0000 UTC"
	date, err := time.ParseInLocation(layout, val, time.UTC)
	if err != nil {
		panic(err)
	}
	// UTC 时间: 2023-01-09 09:54:45.3580496 +0000 UTC
	fmt.Println(date)
	// 本地时间(北京): 2023-01-09 17:54:45
	fmt.Println(date.Local().Format("2006-01-02 15:04:05"))
}

func TestSub(t *testing.T) {
	before := time.Now().Add(-time.Hour * 24)
	fmt.Println(before)
	fmt.Println(time.Since(before).Milliseconds() / 1000)
}

func Test0001(t *testing.T) {
	var a uint64 = 2
	var b uint64 = 4
	c := a - b
	fmt.Println(c)
}
