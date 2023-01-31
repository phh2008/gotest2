package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	const token string = "EEOv2gI2TT5LabDseWEZI3fcYFNZerHb"
	// 当前时间戳（秒级）
	timestamp := time.Now().UnixMilli() / 1000
	// md5(token+时间戳)
	hash := md5.Sum([]byte(fmt.Sprintf("%s%d", token, timestamp)))
	encryptionToken := fmt.Sprintf("%x", hash)
	fmt.Println("  timestampKey:       ", timestamp)
	fmt.Println("  encryptionToken: ", encryptionToken)
}
