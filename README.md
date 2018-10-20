# gofirst
go 也需要自己的运行环境类似于java的JRE， go的运行环境是goroot,
搭建go的开发环境
1.下载go 的开发包
本文是mac下面使用的homebrew 做的安装，安装最新版本的go版本命令如下：
homebrew install go  

2.安装完成后可以通过命令查看go的版本信息：
go version

3.可以通过如下命令查看go的环境变量信息：
go env

结果如下：
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

4.在idea中配置go开发插件：

Preferences--> plugins --> Browse repositories --> go --> go LANGUAGES (下载次数最多的)--> install --> reset idea 

5.配置gopath
可以在Preferences --> Languages & Frameworks --> go 目录下修改goroot gopath 
如3中所示：新建gopath路径，并新建bin、src、pkg三个目录

6.全局配置gopath
在.bash_profile 文件中增加
#go的安装目录
export GOROOT=/usr/local/Cellar/go/1.11.1/libexec
#go 的引入包路径
export GOPATH=/Users/shibinbin/go/go_path
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
并通过source .bash_profile 启用后 通过go env 就能看到新修改的gopath了

 
