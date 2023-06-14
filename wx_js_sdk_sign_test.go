package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestJsSdkSign(t *testing.T) {
	client := resty.New().SetDebug(true)
	appid := "xxx"
	secret := "xxxx"
	// 获取 access_token
	accessUrl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	res, err := client.R().Get(fmt.Sprintf(accessUrl, appid, secret))
	if err != nil {
		panic(err)
	}
	log.Println("access_token-res:", res.String())
	access_token := gjson.Get(res.String(), "access_token").String()
	log.Println("access_token:", access_token)
	// 获取 jsapi_ticket
	ticketUrl := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	res, err = client.R().Get(fmt.Sprintf(ticketUrl, access_token))
	if err != nil {
		panic(err)
	}
	log.Println("ticket-res:", res.String())
	ticket := gjson.Get(res.String(), "ticket").String()
	log.Println("ticket:", ticket)
	// 生成JS-SDK权限验证的签名
	noncestr := RandomString(16)
	timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	url := "http://10.250.221.44:5173"
	sign := Signature(ticket, noncestr, timestamp, url)
	fmt.Println("--------------------------------")
	fmt.Println("appid", appid)
	fmt.Println("noncestr", noncestr)
	fmt.Println("timestamp", timestamp)
	fmt.Println("url", url)
	fmt.Println("sign", sign)
}

func TestNonstr(t *testing.T) {
	fmt.Println(RandomString(16))
}

var char = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var random = rand.New(rand.NewSource(time.Now().Unix()))

func RandomString(length int) string {
	size := len(char)
	b := make([]byte, length)
	for i := range b {
		b[i] = char[random.Intn(size)]
	}
	return string(b)
}

func Signature(jsTicket, noncestr, timestamp, url string) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsTicket, noncestr, timestamp, url)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
