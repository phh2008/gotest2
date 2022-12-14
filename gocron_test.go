package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"testing"
	"time"
)

var YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
var task = func() {
	fmt.Println("task running ... ", time.Now().Format("2006-01-02 15:04:05"))
}

func Test_IsRunning(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	j, _ := s.Every(10).Seconds().Do(func() { time.Sleep(2 * time.Second) })

	fmt.Println(j.IsRunning())

	s.StartAsync()

	time.Sleep(time.Second)
	fmt.Println(j.IsRunning())

	time.Sleep(time.Second)
	s.Stop()

	time.Sleep(1 * time.Second)
	fmt.Println(j.IsRunning())
	// Output:
	// false
	// true
	// false
}

func Test_LastRun(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	job, _ := s.Every(1).Second().Do(task)
	s.StartAsync()
	fmt.Println("Last run:", job.LastRun())
	select {}
}

func Test_LimitRunsTo(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	job, _ := s.Every(1).Second().Do(func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		fmt.Println("aaa")
		var a = 1
		var b = 0
		fmt.Println(">", a/b)
	})
	job.LimitRunsTo(5)
	s.StartBlocking()
}

func Test_RunCount(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	job, _ := s.Every(1).Second().Do(task)
	go func() {
		for {
			fmt.Println("Run count", job.RunCount())
			time.Sleep(time.Second)
		}
	}()
	s.StartBlocking()
}

func Test_ScheduledTime(t *testing.T) {
	s := gocron.NewScheduler(time.Local)
	job, _ := s.Every(1).Day().At("15:17").Do(task)
	fmt.Println(job.ScheduledTime())
	s.StartBlocking()
}

func Test_SingletonMode(t *testing.T) {
	s := gocron.NewScheduler(time.Local)
	s.SingletonModeAll()
	_, _ = s.Every(1).Second().Do(func() {
		fmt.Println("task ...", time.Now().Format(YYYY_MM_DD_HH_MM_SS))
		time.Sleep(time.Second * 3)
	})
	//job.SingletonMode()
	_, _ = s.Every(1).Second().Do(func() {
		fmt.Println("task2 ...", time.Now().Format(YYYY_MM_DD_HH_MM_SS))
		time.Sleep(time.Second * 3)
	})
	s.StartBlocking()
}

func Test_Tag(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	job, _ := s.Every("1s").Do(task)
	job.Tag("tag1", "tag2", "tag3")
	s.StartAsync()
	fmt.Println(job.Tags())
	// Output:
	// [tag1 tag2 tag3]
}

func Test_Untag(t *testing.T) {
	s := gocron.NewScheduler(time.Local)
	job, _ := s.Every("1s").Do(task)

	job.Tag("tag1", "tag2", "tag3")
	s.StartAsync()
	fmt.Println(job.Tags())
	job.Untag("tag2")
	fmt.Println(job.Tags())
	// Output:
	// [tag1 tag2 tag3]
	// [tag1 tag3]
}

func Test_Tag2(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	s.TagsUnique()
	_, _ = s.Every("1s").Tag("foo").Do(task)
	_, err := s.Every("1s").Tag("foo").Do(task)
	if err != nil {
		fmt.Println(err)
	}
	s.StartAsync()
	fmt.Println("job count ", s.Len())

	_, err = s.Every("1s").Tag("foo2").Do(func() { fmt.Println("task 2 ......") })
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("job count ", s.Len())
	select {}
}

func Test_Tag3(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(1).Second().Tag("tag1", "tag0").Do(task)
	_, _ = s.Every(1).Second().Tag("tag2").Do(func() { fmt.Println("task 2...") })
	s.StartAsync()

	time.Sleep(time.Second * 2)
	_ = s.RemoveByTag("tag1")
	fmt.Println(s.Len())

	select {}
}

func Test_WaitForSchedule(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)

	_, _ = s.Every("5s").WaitForSchedule().Do(task)

	_, _ = s.Every("5s").Do(func() { fmt.Println("task 2...") })

	s.StartAsync()
	select {}
}

func Test_SetMaxConcurrentJobs(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	s.SetMaxConcurrentJobs(2, gocron.RescheduleMode)
	_, _ = s.Every(1).Seconds().Do(func() {
		fmt.Println("This will run once every 5 seconds even though it is scheduled every second because maximum concurrent job limit is set.")
		time.Sleep(5 * time.Second)
	})
	_, _ = s.Every("1s").Do(task)
	s.StartAsync()
	select {}
}
