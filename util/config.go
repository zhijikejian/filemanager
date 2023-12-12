package util

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// route监听地址
type ListenAddr struct {
	Ip   string `ini:"ip"`
	Port string `ini:"port"`
}

var Listenaddr ListenAddr

type Folder struct {
	Name   string `ini:"name"`
	Path   string `ini:"path"`
	Active bool   `ini:"active"`
}

var Folders = []Folder{}

// 初始化，解析配置文件
func Loadconfig() bool {
	fmt.Println("reading filemanager.ini ...")
	cfg, err := ini.Load("filemanager.ini")
	if err != nil {
		fmt.Println("read filemanager.ini err:", err)
		return false
	}
	for _, key := range cfg.SectionStrings() {
		if key == "DEFAULT" {
		} else if key == "common" {
			var listenaddr = ListenAddr{
				Ip:   "",
				Port: "8989",
			}
			err := cfg.Section(key).MapTo(&listenaddr)
			if err != nil {
				fmt.Println("mapto listenaddr err", err)
			} else {
				Listenaddr = listenaddr
			}
		} else {
			var folder = Folder{
				Name:   key,
				Path:   "",
				Active: true,
			}
			err := cfg.Section(key).MapTo(&folder)
			if err != nil {
				fmt.Println("mapto folder err", err)
			} else {
				if folder.Active {
					Folders = append(Folders, folder)
				}
			}
		}
	}
	return true
}
func init() {
	Loadconfig()
}
func main() {
	// 测试
	// Loadconfig()
	fmt.Println("Listenaddr", Listenaddr)
	fmt.Println("Folders", Folders)
}
