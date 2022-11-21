package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtx1(t *testing.T) {
	jobs := map[string]context.CancelFunc{}
	// add job
	ctx, cancel := context.WithCancel(context.Background())
	jobs["1000"] = cancel
	go doTask(ctx, "1000")
	ctx2, cancel2 := context.WithCancel(context.Background())
	jobs["1002"] = cancel2
	go doTask(ctx2, "1002")

	time.Sleep(3 * time.Second)
	if c, ok := jobs["1000"]; ok {
		c() // 取消任务
	}
	select {}
}

func doTask(ctx context.Context, args string) {
	fmt.Println("task running ......, args: ", args)
	select {
	case <-ctx.Done():
		fmt.Println("job closed ...... args: ", args)
	}
}
