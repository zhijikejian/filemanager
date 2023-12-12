package routers

import (
	"file_views"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	files := e.Group("/filedata")
	{
		files.POST("/filelist", file_views.Filelist_view)
		files.POST("/folder", file_views.Folder_view)
		files.POST("/upload", file_views.Upload_view)
		files.GET("/download", file_views.Download_view)
		files.POST("/delete", file_views.Delete_view)
		files.GET("/login", file_views.Login_view)
	}
}
