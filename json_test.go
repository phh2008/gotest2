package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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
	var json = `{"dmx_info":{"active":true,"from_device":"sacn","DMX Protocol Version":"4.4","DMX Address":68,"DMX Mode Number":3,"DMX Mode Name":"CCT & HSI 8Bit","footprint":18,"signal_lost_name":"Hold","signal_lost_nb":1,"signal_lost_state":false,"extended_color":true,"rdm_state":true,"dmx_stat":{"signal_cnt":0,"refresh_time":0,"frame_cnt":0,"packet_cnt":7589,"last_process_time":83,"xlr":{"rx_active":false,"rx_cnt":0,"tx_cnt":81354,"break_cnt":0},"artnet":{"rx_active":false,"rx_cnt":23818,"tx_cnt":515},"sacn":{"rx_active":true,"rx_cnt":7589,"tx_cnt":0}},"rdm":{"uid":"20B960CC771F","xlr":{"mute":false,"Active":false,"rx_time":0,"tx_cnt":0,"rx_cnt":0,"queue_cnt":1,"queue_error_cnt":0,"csum_error_cnt":0}},"light_data":{"Intensity":11.76,"Intensity DMX":30,"CCT 1":3195.29,"CCT 1 DMX":14,"G/M 1":0.0,"G/M 1 DMX":128,"X-Fade":0.0,"X-Fade DMX":0,"Hue 2":0.0,"Hue 2 DMX":0,"Sat 2":0.0,"Sat 2 DMX":0,"DMX Fan Mode":"None","DMX Fan Mode DMX":0,"Preset":"None","Preset DMX":0,"Strobe":0.0,"Strobe DMX":0,"+/- Warmer":0.0,"+/- Warmer DMX":0,"+/- Sat":0.0,"+/- Sat DMX":0,"+/- Red":0.0,"+/- Red DMX":0,"+/- Green":0.0,"+/- Green DMX":0,"+/- Blue":0.0,"+/- Blue DMX":0,"+/- Cyan":0.0,"+/- Cyan DMX":0,"+/- Magenta":0.0,"+/- Magenta DMX":0,"+/- Yellow":0.0,"+/- Yellow DMX":0},"raw_data":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,30,14,128,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]}}`
	value := gjson.Get(json, "dmx_info.light_data.DMX Fan Mode")
	log.Println(value.String())
}

func TestGjson03(t *testing.T) {
	var json = `{"ntcLightEngine":[{"actual":22.1777,"min":20.344,"max":22.2653},{"actual":21.9151,"min":20.2569,"max":21.9151},{"actual":22.0026,"min":20.4311,"max":22.0026}],"ntcPwmDrv":[{"actual":24.7335,"min":21.0414,"max":24.7335},{"actual":24.3788,"min":20.8669,"max":24.4674},{"actual":24.9111,"min":21.2159,"max":25.0},{"actual":25.0,"min":21.3033,"max":25.089}],"ntcMain":{"actual":25.6015,"min":20.7144,"max":25.6238},"ntcBoost":{"actual":23.82,"min":21.32,"max":23.84},"supply":{"U-In-Main":49.52,"U-In-Bat":0.28,"I/O-Bat":false,"U-Out":49.53,"I-In":0.35,"P-In":17.332,"U-12V":12.3352,"U-5V":5.0918},"counter":{"fixture":1817836,"ledengine":280932,"fan":1815098},"fan":{"fan1":2173.91}}`
	path := "@this"
	value := gjson.Get(json, path)
	//log.Println(value)
	value.ForEach(func(key, value gjson.Result) bool {
		log.Println(key, " :: ", value)
		return true
	})
}

func TestGjson04(t *testing.T) {
	var json = `{
"dmx_info": {
	"active": false,
	"from_device": "no supp",
	"DMX Protocol Version": "4.4",
	"DMX Address": 68,
	"DMX Mode Number": 3,
	"DMX Mode Name": "CCT & HSI 8Bit",
	"footprint": 18,
	"signal_lost_name": "Hold",
	"signal_lost_nb": 1,
	"signal_lost_state": false,
	"extended_color": true,
	"rdm_state": true,
	"dmx_stat": {
		"signal_cnt": 0,
		"refresh_time": 0,
		"frame_cnt": 0,
		"packet_cnt": 0,
		"last_process_time": 0,
		"xlr": {
			"rx_active": false,
			"rx_cnt": 0,
			"tx_cnt": 0,
			"break_cnt": 0},
		"artnet": {
			"rx_active": false,
			"rx_cnt": 2,
			"tx_cnt": 2},
		"sacn": {
			"rx_active": false,
			"rx_cnt": 0,
			"tx_cnt": 0}},
	"rdm": {
		"uid": "20B960CC771F",
		"xlr": {
			"mute": false,
			"Active": false,
			"rx_time": 0,
			"tx_cnt": 0,
			"rx_cnt": 0,
			"queue_cnt": 0,
			"queue_error_cnt": 0,
			"csum_error_cnt": 0}},
	"raw_data": [
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
		0,0,0]
	}}`
	path := "@this"
	value := gjson.Get(json, path)
	//log.Println(value)
	value.ForEach(func(key, value gjson.Result) bool {
		log.Println(key, " :: ", value)
		return true
	})
}

func TestSjson01(t *testing.T) {
	json := `[{"first":"a1","last":"a2"},{"first":"b1","last":"b2"}]`
	arr := gjson.Get(json, "@this").Array()
	value, _ := sjson.SetRaw("[]", "0", `{"first":"c1","last":"c2"}`)
	for _, v := range arr {
		value, _ = sjson.SetRaw(value, "-1", v.Raw)
	}
	println(value)
}

func TestGjson05(t *testing.T) {
	json := `{"EventType":"FileDeleted","FileDeleteEvent":{"FileIdSet":["243791581461656546","11111122222"],"FileDeleteResultInfo":[{"FileId":"243791581461656546","DeleteParts":[]}]}}`
	result := gjson.Get(json, "@this")
	arr := result.Get("FileDeleteEvent.FileIdSet").Array()
	for _, v := range arr {
		fmt.Println(v.String())
	}
}
