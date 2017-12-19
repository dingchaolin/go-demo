# go-demo

### mac go环境配置
- brew install go

### 查看版本
- go version

### 获取依赖
- go get 
- -u 让命令利用网络来更新已有代码包及其依赖包。默认情况下，该命令只会从网络上下载本地不存在的代码包，而不会更新已有的代码包。
- -x 显示go get命令执行过程中所使用的所有命令
- -v 标记意味着会打印出被构建的代码包的名字
- go get -v -u -x github.com/hyper-carrot/go_lib

# 什么是go
- go是一门并发支持，垃圾回收的编译系统编程语言，圣在创造一门具有在静态编译语言的高性能和动态语言的高效开发之间拥有一个良好的平衡点。

# go的主要特点
- 类型安全和内存安全
- 以非常直观和极低代价的方案实现高并发
- 高效的垃圾回收机制
- 快速编译
- 为多核计算机提供性能提升的方案
- utf8编码支持

# 开源网站
- go-wiki

# 环境变量
```
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/chaolinding/go"//工作目录
GORACE=""
GOROOT="/usr/local/Cellar/go/1.9.2/libexec"
GOTOOLDIR="/usr/local/Cellar/go/1.9.2/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/4t/vw_tmxb51g7_5jq9_rf98kl80000gn/T/go-build259991128=/tmp/go-build -gno-record-gcc-switches -fno-common"
CXX="clang++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
```
# gopath
- 有3个目录
- bin 存放可执行文件的目录
- pkg 存放编译生成的.a静态库
- src 存放源码

# go命令
- go get 安装远程包
- go run 执行代码
- go build 编译
- go fmt 格式化源码
- go install 编译包文件 编译主程序 编译好的包文件放到pkg下，编译好的主程序放到bin下
- go test 运行测试文件 _test.go 结尾的文件都是测试文件 会自动运行_test.go结尾的文件
- go doc 查看包 go doc fmt   或者  godoc fmt PrintLln
- 建立本地官网查看文档 godoc -http=:8888 在本地就可以以localhost:8888访问官网文档了

# go install
- 在操作系统名称_cpu架构文件夹下生成.a 文件

# go build
- 生成可执行文件

# go run
- go run main.go 生成可执行文件

 


