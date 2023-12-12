package file_views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"util"

	"github.com/gin-gonic/gin"
)

var folderlist = gin.H{}

func Loadfolderlist() {
	fmt.Println("load folderlist ...")
	folderlist = loadfolder()
}

var logtime = util.NewSyncMap() // 记录所有账号的存活时间

func loglive(user string) bool {
	var timenow = time.Now().Unix()
	if usertime, ok := logtime.Load(user); ok {
		if timenow-usertime.(int64) <= 300 { // 5min
			logtime.Store(user, timenow)
			return true
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}

func Login_view(c *gin.Context) {
	fmt.Println("Login_view")
}

func Filelist_view(c *gin.Context) {
	fmt.Println("Filelist_view")
	var p = map[string]string{}
	b, _ := c.GetRawData()
	_ = json.Unmarshal(b, &p)
	if p["refresh"] == "refresh" {
		folderlist = loadfolder()
	}
	var folderjson, err = json.Marshal(folderlist)
	if err != nil {
		folderjson = []byte("{}")
	}
	resjson := gin.H{"user": p["user"], "code": "666", "msg": "ok", "data": string(folderjson)}
	c.JSON(http.StatusOK, resjson)
}

func Folder_view(c *gin.Context) {
	fmt.Println("Folder_view")
	var p = map[string]string{}
	b, _ := c.GetRawData()
	_ = json.Unmarshal(b, &p)
	for _, folder := range util.Folders {
		if strings.HasPrefix(p["path"], folder.Name) {
			p["path"] = strings.Replace(p["path"], folder.Name, folder.Path, 1)
			break
		}
		if strings.HasPrefix(p["path"], "/"+folder.Name) {
			p["path"] = strings.Replace(p["path"], "/"+folder.Name, folder.Path, 1)
			break
		}
	}
	p["path"] = strings.Replace(p["path"], "\\", "/", -1)
	var filesjson, err = json.Marshal(loadfiles(p["path"]))
	if err != nil {
		filesjson = []byte("{\"isdir\":true,\"folder\":[]}") // gin.H{"isdir": true, "folder": gin.H{}}
	}
	resjson := gin.H{"user": p["user"], "code": "666", "msg": "ok", "data": string(filesjson)}
	c.JSON(http.StatusOK, resjson)
}

func Upload_view(c *gin.Context) {
	fmt.Println("Upload_view")
	var user = c.Query("user")
	var path = c.Query("folder")
	fmt.Println("user", user, "path", path)
	for _, folder := range util.Folders {
		if strings.HasPrefix(path, folder.Name) {
			path = strings.Replace(path, folder.Name, folder.Path, 1)
			break
		}
		if strings.HasPrefix(path, "/"+folder.Name) {
			path = strings.Replace(path, "/"+folder.Name, folder.Path, 1)
			break
		}
	}
	path = strings.Replace(path, "\\", "/", -1)
	var msg = "ok"
	var code = "666"
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		// 上传文件到指定的目录
		var path = path + "/" + file.Filename
		fmt.Println("MultipartForm path", path)
		c.SaveUploadedFile(file, path)
	}
	resjson := gin.H{"user": user, "code": code, "msg": msg, "data": path}
	c.JSON(http.StatusOK, resjson)
}

func Download_view(c *gin.Context) {
	fmt.Println("Download_view")
}

func Delete_view(c *gin.Context) {
	fmt.Println("Delete_view")
	var p = map[string]string{}
	b, _ := c.GetRawData()
	_ = json.Unmarshal(b, &p)
	for _, folder := range util.Folders {
		if strings.HasPrefix(p["path"], folder.Name) {
			p["path"] = strings.Replace(p["path"], folder.Name, folder.Path, 1)
			break
		}
		if strings.HasPrefix(p["path"], "/"+folder.Name) {
			p["path"] = strings.Replace(p["path"], "/"+folder.Name, folder.Path, 1)
			break
		}
	}
	p["path"] = strings.Replace(p["path"], "\\", "/", -1)
	// 删除文件
	var msg = "ok"
	var code = "666"
	err := os.Remove(p["path"])
	if err != nil {
		fmt.Println("os.Remove", p["path"], "err", err)
		msg = "fail"
		code = "000"
	}
	resjson := gin.H{"user": p["user"], "code": code, "msg": msg, "data": p["path"]}
	c.JSON(http.StatusOK, resjson)
}

func childfolder(path string) gin.H {
	var childres = gin.H{}
	f, err := os.Open(path)
	if err != nil {
		return childres
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return childres
	}
	for _, file := range files {
		if file.IsDir() {
			childres[file.Name()] = childfolder(path + "/" + file.Name())
		}
	}
	return childres
}

func loadfolder() gin.H {
	var res = gin.H{}
	for _, folder := range util.Folders {
		s, err := os.Stat(folder.Path)
		if err != nil {
			fmt.Println("Stat err", err)
			continue
		}
		if !s.IsDir() {
			fmt.Println(folder.Path, "不是文件夹")
			continue
		}
		res[folder.Name] = childfolder(folder.Path)
	}
	return res
}

func loadfiles(path string) gin.H {
	var res = gin.H{"isdir": true, "folder": gin.H{}}
	s, err := os.Stat(path)
	if err != nil {
		fmt.Println("Stat err", err)
		return res
	}
	if s.IsDir() {
		f, err := os.Open(path)
		if err != nil {
			return res
		}
		files, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return res
		}
		var childres = []gin.H{}
		for _, file := range files {
			childres = append(childres, gin.H{"name": file.Name(), "isdir": file.IsDir(), "size": file.Size(), "modtime": file.ModTime()})
		}
		res = gin.H{"isdir": true, "folder": childres}
	} else {
		res = gin.H{"isdir": false, "folder": gin.H{}}
	}
	return res
}
