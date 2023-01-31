package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestFileAppend(t *testing.T) {
	// 文件存放路径
	path := "D:\\codespace\\golang_test\\go-test2\\xxx\\"
	// 保存文件名
	filename := "test_file_02.log"

	// 目录是否存在，不存在就创建目录
	if !IsExist(path) {
		CreateDir(path)
	}
	// 文件不存在自动创建，内容写入文件未尾
	file, err := os.OpenFile(path+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
	if err != nil {
		// 打开文件错误处理
		log.Panic(err)
	}
	// 及时关闭file句柄
	defer file.Close()
	// 写入文件，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		// 写入文件内容
		write.WriteString(" hello，你好。222  \n")
	}
	// Flush 将缓存的文件真正写入到文件中
	write.Flush()
}

// IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// CreateDir  文件夹创建
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}
