# filemanager

使用golang-gin编写的 **web版文件管理工具**，web框架是gin



## 配置文件

filemanager.ini

```ini
[common] ; 本地配置
ip=127.0.0.1 ; ip地址，必填项
port=8989 ; 端口号，必填项

[filefolder] ; 共享的一个目录名称
path=D:\filefolder ; 路径，绝对路径
active=true ; 是否生效

[xunlei] ; 共享的一个目录名称
path=X:\迅雷下载
active=true

; 共享的其他目录名称
```

> [common] 是web配置，需要设置ip和port端口号，ip设置为0.0.0.0则为外网可访问
>
> [filefolder] 与 [xunlei] 是配置的两个文件路径名，在管理页面左侧目录树中显示为最高路径名
>
> path 参数则是实际的本地硬盘完整目录路径，其子目录会显示在最高路径名下
>
> active 参数标识这个路径是否开启，true 为开启，false 为不开启

注意：[common] 标签只能有一个，[filefolder] 与 [xunlei] 这种表示路径名额标签则可以有多个，但是名称不能重复



## 使用

```
./filemanager
```



## 开机启动

建议使用filemanager.sh或filemanager.cmd来使用

linux filemanager.sh

```sh
cd /home/xiao/my/filemanager
./filemanager
```

windows filemanager.cmd

```
cd /d D:\my\filemanager
filemanager.exe
```

自启动：

linux可选择新建service来自启动

> 自行查询service文件新建和使用方法

windows可选择将filemanager.cmd的快捷方式放到启动文件夹

> win+R，输入shell:startup回车，就是启动文件夹



## 适用场景

- 如果不想用ftp来管理远端机文件，可以使用filemanager工具来做简单文件管理

- 配合vnet虚拟组网工具使用更香甜



## 使用的其他外部库

- gopkg.in/ini.v1      用来读取ini配置文件
- github.com/gin-gonic/gin      filemanager服务后端
- vue      filemanager前端处理框架
- element-plus      filemanager前端组件框架
- axios      filemanager前端ajax框架



## 注意事项

- 下载并不是通用支持，有些浏览器会跳转播放

- filemanager无身份验证支持

- 配合vnet-v1版本使用情况下，播放在线视频可能导致vnet端其他服务卡顿，这是由于vnet-v1版本采用单链路与route端通信的缘故。vnet-v2版本不存在这个问题

    