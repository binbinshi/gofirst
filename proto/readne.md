要生成*pb.go文件，需要执行protoc命令，先要下载protobuf
1.mac 下依然使用homebrew来安装

```
brew install protobuf
```
2.需要安装protoc-gen-go 执行protoc命令
```
go get github.com/golang/protobuf/protoc-gen-go
```
3.执行命令
需要进入xxx.proto文件所在的目录进行执行以下命令
```
protoc -I. binbin.proto --go_out=plugins=grpc:/Users/shibinbin/go/go_path/src/gofirstdemo/proto

protoc -I. binbin.proto --go_out=plugins=grpc:../proto
```
会生成xxx.pb.go文件