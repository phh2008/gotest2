package main

import (
	"context"
	"fmt"
	iotClient "github.com/apache/iotdb-client-go/client"
	"github.com/apache/iotdb-client-go/rpc"
	pool "github.com/jolestar/go-commons-pool/v2"
	"log"
	"math/rand"
	"testing"
	"time"
)

var (
	host     string = "127.0.0.1"
	port     string = "6667"
	user     string = "root"
	password string = "root"
)

var config *iotClient.Config = &iotClient.Config{
	Host:     host,
	Port:     port,
	UserName: user,
	Password: password,
}

// var session iotClient.Session
var sessionPool *pool.ObjectPool

func init() {
	factory := pool.NewPooledObjectFactorySimple(
		func(context.Context) (interface{}, error) {
			session := iotClient.NewSession(config)
			err := session.Open(false, 0)
			return &session, err
		})
	ctx := context.Background()
	poolConfig := pool.NewDefaultPoolConfig()
	poolConfig.MaxTotal = 500
	poolConfig.MinIdle = 20
	poolConfig.MinEvictableIdleTime = -1
	sessionPool = pool.NewObjectPool(ctx, factory, poolConfig)
}

func TestSetStorageGroup(t *testing.T) {
	obj, err := sessionPool.BorrowObject(context.Background())
	if err != nil {
		log.Panic(err)
	}
	session := obj.(*iotClient.Session)
	status, err := session.SetStorageGroup("root.ln2")
	checkError(status, err)
}

func TestCreateTimeSeries(t *testing.T) {
	obj, err := sessionPool.BorrowObject(context.Background())
	if err != nil {
		log.Panic(err)
	}
	session := obj.(*iotClient.Session)
	checkError(session.CreateTimeseries("root.ln.g1.d1.temperature", iotClient.FLOAT, iotClient.PLAIN, iotClient.SNAPPY, nil, nil))
	checkError(session.CreateTimeseries("root.ln.g1.d1.status", iotClient.BOOLEAN, iotClient.PLAIN, iotClient.SNAPPY, nil, nil))
}

func TestDeleteTimeSeries(t *testing.T) {
	obj, err := sessionPool.BorrowObject(context.Background())
	if err != nil {
		log.Panic(err)
	}
	session := obj.(*iotClient.Session)
	checkError(session.DeleteTimeseries([]string{"root.ln.g1.d1"}))
}

func TestInsert(t *testing.T) {
	obj, err := sessionPool.BorrowObject(context.Background())
	if err != nil {
		log.Panic(err)
	}
	session := obj.(*iotClient.Session)
	checkError(session.InsertStringRecord("root.ln.g1.d1", []string{"temperature", "status"}, []string{"37.6", "false"}, time.Now().UnixMilli()))
}

func TestInsertMore(t *testing.T) {
	obj, err := sessionPool.BorrowObject(context.Background())
	if err != nil {
		log.Panic(err)
	}
	session := obj.(*iotClient.Session)
	date, _ := time.Parse("2006-01-02 15:04:05", "2022-08-01 00:00:00")
	rand.Seed(time.Now().UnixNano())
	names := []string{"tom", "jack", "lili", "lucy", "张三丰", "李四", "王五"}
	namesSize := len(names)
	for i := 1; i < 1000000; i++ {
		dt := date.Add(time.Second * time.Duration(i))
		n := rand.Float64()*(39-35) + 35
		tp := fmt.Sprintf("%.1f", n)
		status := "false"
		if rand.Intn(2) == 1 {
			status = "true"
		}
		name := names[rand.Intn(namesSize)]
		fmt.Println("i=", i, "time: ", dt, " status: ", status, " temperature: ", tp, " name: ", name)
		checkError(session.InsertStringRecord("root.ln.g2.d1", []string{"temperature", "status", "name"}, []string{tp, status, name}, dt.UnixMilli()))
	}
}

func TestName(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Float64()*(39-35) + 35
	fmt.Println(fmt.Sprintf("%.1f", n))
	fmt.Println(rand.Intn(2))
	names := []string{"tom", "jack", "lili", "lucy", "张三丰", "李四", "王五"}
	fmt.Println(names[rand.Intn(len(names))])
}

func TestQuery(t *testing.T) {
	obj, err := sessionPool.BorrowObject(context.Background())
	if err != nil {
		log.Panic(err)
	}
	session := obj.(*iotClient.Session)
	var timeout int64 = 1000
	//if ds, err := session.ExecuteQueryStatement("select last status,name,temperature from root.ln.g1.d3 order by timeseries", &timeout); err == nil {
	if ds, err := session.ExecuteQueryStatement("select status,name,temperature from root.ln.g1.d3 order by time desc limit 1", &timeout); err == nil {
		printDevice1(ds)
		ds.Close()
	} else {
		log.Fatal(err)
	}
}

func printDevice1(sds *iotClient.SessionDataSet) {
	showTimestamp := !sds.IsIgnoreTimeStamp()
	if showTimestamp {
		fmt.Print("Time\t\t\t\t")
	}
	for _, columnName := range sds.GetColumnNames() {
		fmt.Printf("%s\t", columnName)
	}
	fmt.Println()

	for next, err := sds.Next(); err == nil && next; next, err = sds.Next() {
		if showTimestamp {
			fmt.Printf("%s\t", sds.GetText(iotClient.TimestampColumnName))
		}
		var name string
		var value string
		var dataType string

		// All of iotdb datatypes can be scan into string variables
		if err := sds.Scan(&name, &value, &dataType); err != nil {
			log.Fatal(err)
		}
		whitespace := "\t\t"
		fmt.Printf("%v%s", name, whitespace)
		fmt.Printf("%v%s", value, whitespace)
		fmt.Printf("%v%s", dataType, whitespace)
		fmt.Println("\n............................................")

		if record, err := sds.GetRowRecord(); err == nil {
			for _, field := range record.GetFields() {
				v := field.GetValue()
				if field.IsNull() {
					v = "null"
				}
				fmt.Printf("%v\t\t", v)
			}
			fmt.Println()
		}
		fmt.Println("----------------------------------------------")
	}
}

func checkError(status *rpc.TSStatus, err error) {
	if err != nil {
		log.Fatal("error: ", err)
	}

	if status != nil {
		if err = iotClient.VerifySuccess(status); err != nil {
			log.Println("VerifySuccess: ", err)
		}
	}
}
