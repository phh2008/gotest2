package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFilepath(t *testing.T) {
	// Dir：获取path中最后一个分隔符之前的部分(不包含分隔符)
	fmt.Println("Dir：", filepath.Dir("D:file/upload/image/abc.jpg"))
	fmt.Println("Dir：", filepath.Dir("D:file/upload/image/"))
	fmt.Println("Dir：", filepath.Dir("D:file\\upload\\image\\"))
	fmt.Println("Dir：", filepath.Dir("D:file\\upload\\image"))

	// Base：获取path中最后一个分隔符之后的部分(不包含分隔符)
	fmt.Println("Base：", filepath.Base("D:file/upload/image/abc.jpg"))
	fmt.Println("Base：", filepath.Base("D:file/upload/image/"))
	fmt.Println("Base：", filepath.Base("D:file/upload\\image"))

	// Split：获取path中最后一个分隔符前后的两部分，
	fmt.Println("Split：")
	fmt.Println(filepath.Split("D:file/upload/image/abc.jpg"))
	fmt.Println(filepath.Split("D:file\\upload\\image\\abc.jpg"))
	fmt.Println(filepath.Split("/upload/image/"))
	fmt.Println(filepath.Split("/upload/image"))

	// Ext：获取路径字符串中的文件扩展名
	fmt.Println("Ext：", filepath.Ext("D:file/upload/image/abc.jpg"))
	fmt.Println("Ext：", filepath.Ext("D:file\\upload\\image\\abc.jpg"))
	fmt.Println("Ext：", filepath.Ext("D:file\\upload\\image\\"))
	fmt.Println("Ext：", filepath.Ext("./file/upload/image"))

	fmt.Println(filepath.Join("/upload/log", "1002", "device01", "aaaaaaa.log"))

	fmt.Println("\\\\转换成/", filepath.ToSlash("D:file\\upload\\image\\abc.jpg"))

}
