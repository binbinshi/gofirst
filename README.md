# gofirst
go 也需要自己的运行环境类似于java的JRE， go的运行环境是goroot,
搭建go的开发环境
1.下载go 的开发包
本文是mac下面使用的homebrew 做的安装，安装最新版本的go版本命令如下：
```
brew install go  
```
2.安装完成后可以通过命令查看go的版本信息：
```
go version
```
3.可以通过如下命令查看go的环境变量信息：
```
go env
```
结果如下：
```
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/shibinbin/Library/Caches/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/shibinbin/go/go_path"
GOPROXY=""
GORACE=""
GOROOT="/usr/local/Cellar/go/1.11.1/libexec"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.11.1/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/gw/r5bmb0y51r5dg_5jgc39gp100000gp/T/go-build375474522=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

4.在idea中配置go开发插件：
```
Preferences--> plugins --> Browse repositories --> go --> go LANGUAGES (下载次数最多的)--> install --> reset idea 
```
5.配置gopath
gopath 即使你的工作目录：
bin (存放编译后生成的可执行文件)
pkg（存放编译后生成的包文件）
src（存放项目源码）
可以在Preferences --> Languages & Frameworks --> go 目录下修改goroot gopath 
如3中所示：新建gopath路径，并新建bin、src、pkg三个目录

6.全局配置gopath
在.bash_profile 文件中增加
#go的安装目录
```
export GOROOT=/usr/local/Cellar/go/1.11.1/libexec
#go 的引入包路径
export GOPATH=/Users/shibinbin/go/go_path
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
并通过
```
source .bash_profile
```
启用后 通过go env 就能看到新修改的gopath了

go的常用命令 
go get ：获取远程包（需要安装git） 如下例子，下载后可以在go_path路径下的pkg下查看包文件和src下查看源码；
pkg下是.a结尾的文件，src下则是具体的.go文件源码
```
go get github.com/labstack/echo
```
go run ： 直接运行程序
运行当前程序，命令如下：
```
go run Hello.go
```
当程序中有import 但是未使用的包，程序会报错，如下所示：
```
# command-line-arguments
./Hello.go:5:2: imported and not used: "github.com/labstack/echo"

```
另外\n在末尾要有，不然会出现%
```
Hello Go !% 
```
带\n的输出结果是：
```
Hello Go !
```
go build ：测试编译，检查是否有错误
命令如下所示：
```
go build Hello.go
```
当代码有错误时，编译不通过，并会提示，例如：
```
can't load package: package main: 
Hello.go:4:2: string literal not terminated

```
go fmt ：格式化源代码（部分IDE在保存时自动调用）
格式化代码，go的首行缩进是一个tab的倍数，在命令行中如下执行即可
```
go fmt Hello.go
```
go install ： 编译包文件并编译整个程序
需要先get到你的工作空间才能进行install操作
```
go install github.com/stretchr/testify/assert
```
go test ：运行测试文件
go doc ： 查看文档
```
go doc fmt 
```
查看某个包里的某个函数
```
godoc fmt Print
```
构建一个本地官网，查看帮助文档
```
godoc -http=:8080
```





 
