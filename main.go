package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"file_views"
	"routers"
	"util"

	"github.com/gin-gonic/gin"
)

func init() {
	// gin.SetMode(gin.ReleaseMode)
	file_views.Loadfolderlist()
}
func main() {
	// 创建记录日志的文件
	f, _ := os.OpenFile("logs/ginweb.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer f.Close() // 好像没啥用，ctrl+c退出
	r := gin.New()
	// 日志
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: io.MultiWriter(os.Stdout, f),
		Formatter: func(param gin.LogFormatterParams) string {
			var statusColor, methodColor, resetColor string
			if param.IsOutputColor() {
				statusColor = param.StatusCodeColor()
				methodColor = param.MethodColor()
				resetColor = param.ResetColor()
			}
			// 处理请求超过一分钟，将处理请求时间对秒取整
			if param.Latency > time.Minute {
				// Truncate in a golang < 1.8 safe way
				param.Latency = param.Latency - param.Latency%time.Second
			}
			return fmt.Sprintf("[GIN] %s [%s] |%s %3d %s| %13v |%s %-7s %s %s | %s \"%s\" %s \n",
				param.ClientIP,
				param.TimeStamp.Format("2006-01-02 15:04:05"),
				statusColor, param.StatusCode, resetColor,
				param.Latency,
				methodColor, param.Method, resetColor,
				param.Path,
				param.Request.Proto,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
	}))
	r.Use(gin.Recovery())
	r.StaticFile("/", "./templates/file/file.html")
	r.StaticFS("static", http.Dir("./static"))
	r.StaticFS("file", http.Dir("./templates/file"))
	// 收录配置中的文件路径
	for _, folder := range util.Folders {
		r.StaticFS(folder.Name, http.Dir(folder.Path))
	}
	routers.Routers(r)
	// var listenaddr = "127.0.0.1:8989"
	var listenaddr = util.Listenaddr.Ip + ":" + util.Listenaddr.Port
	if err := r.Run(listenaddr); err != nil {
		fmt.Println("ginstaticserver 启动失败。 err:%v\n", err)
	}
}
