package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/tidwall/gjson"
	"log"
	"testing"
	"time"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type Device struct {
	Id             string    `json:"id"`
	StudioId       int       `json:"studioId"`
	DeviceSnmpPort int32     `json:"deviceSnmpPort"`
	DeviceStatus   string    `json:"deviceStatus"`
	CreateTime     time.Time `json:"createTime"`
}

func TestMarshal(t *testing.T) {
	d := Device{
		Id:             "3",
		StudioId:       1002,
		DeviceSnmpPort: 161,
		DeviceStatus:   "0",
		CreateTime:     time.Now(),
	}
	result := Result[Device]{Code: 0, Msg: "Success", Data: d}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	j, err := json.Marshal(&result)
	if err != nil {
		panic(err)
	}
	log.Println(string(j))
}

func TestUnmarshal(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var j = "{\"code\":0,\"msg\":\"Success\",\"data\":{\"id\":3,\"studioId\":\"1002\",\"deviceSnmpPort\":161,\"deviceStatus\":\"0\",\"createTime\":\"2023-01-06T09:32:03.894+08:00\"}}"
	var result Result[Device]
	err := json.Unmarshal([]byte(j), &result)
	if err != nil {
		panic(err)
	}
	log.Printf("%#v\n", result)
}

func TestGjson01(t *testing.T) {
	var json = "{\"code\":0,\"msg\":\"Success\",\"data\":{\"id\":\"3\",\"studioId\":1002,\"deviceSnmpPort\":161,\"deviceStatus\":\"0\",\"createTime\":\"2023-01-06T09:32:03.894+08:00\"}}"
	value := gjson.Get(json, "data.studioId")
	println(value.String())
}

func TestGjson02(t *testing.T) {
}
