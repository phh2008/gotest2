package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const signKey = "sign"
const timestampKey = "timestamp"
const anonKey = "anon"
const signToken = "46864decfb9211ed8a0900ff36392f40"

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var innerRand = rand.New(rand.NewSource(time.Now().UnixMilli()))

type mapFlag map[string]string

func (f mapFlag) String() string {
	return fmt.Sprintf("%v", map[string]string(f))
}

func (f mapFlag) Set(value string) error {
	split := strings.SplitN(value, "=", 2)
	if len(split) < 2 {
		return errors.New("参数错误")
	}
	f[split[0]] = split[1]
	return nil
}

func main() {
	var mf = mapFlag{}
	flag.Var(&mf, "q", "q list,for example: -q key1=value1 -q key2=value2")
	flag.Parse()
	var data = map[string]interface{}{
		anonKey:      Random(16),
		timestampKey: strconv.FormatInt(time.Now().Unix(), 10),
	}
	for k, v := range mf {
		data[k] = v
	}
	sign := createSign(data)
	data[signKey] = sign
	for k, v := range data {
		fmt.Printf("%25s : %s \n", k, v)
	}
	bt, _ := json.Marshal(&data)
	fmt.Println(string(bt))
}

// createSign 生成签名串
func createSign(data map[string]interface{}) string {
	var keys []string
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for _, k := range keys {
		if k == signKey {
			continue
		}
		val := data[k]
		if val != nil {
			valType := reflect.ValueOf(val)
			kind := valType.Type().Kind()
			if kind == reflect.Struct || kind == reflect.Slice || kind == reflect.Map {
				continue
			}
		}
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(cast.ToString(val))
		sb.WriteString("&")
	}
	values := strings.TrimSuffix(sb.String(), "&") + signToken
	sign := md5.Sum([]byte(values))
	return hex.EncodeToString(sign[:])
}

func Random(length int) string {
	size := len(chars)
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		idx := innerRand.Intn(size)
		result[i] = chars[idx]
	}
	return string(result)
}
