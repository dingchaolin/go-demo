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

# 内置关键字
- break default func interface select case defer go map struct
- chan else goto package switch const fallthrough if range type
- continue fir import return var

# go程序结构
- go程序是通过package来组织的
- package 必须放在非注释的第一行
- 只有package名称为main的包可以包含main函数
- 一个可执行程序有且仅有一个main包
- 通过import关键字来导入其他非main包
- 通过const关键字来进行常量的定义
- 通过在函数体外部使用var关键字进行全局变量的声明与赋值
- 通过type关键字进行结果struct或者接口interface的声明
- 通过func关键字来进行函数的声明

# 可见性规则
- go语言中，使用大小写来决定该常量 变量 类型 接口 结构 或者 函数 是否可以被外部包所调用
- 根据约定 函数首字母小写即为private， 函数名首字母为大写 为 public

# 基本类型
- bool true false 只能使用true false 不能使用数字来代替true 或者 false 1字节
- int/ uint 32位 或者 64位
- int8 -128~127  uint8 0~255 1字节
- byte uint8的别名
- int16/uint16 2字节
- int32(rune)/uint32 4字节
- rune 用于做unicode相关的处理
- int64/uint64 8字节
- float32/float64 4/8字节 精确7/15位小数
- 复数 complex64/complex128  8/16字节
- uintptr  根据平台区分 32位还是64位  保存指针
- 其他值类型 array struct string
- 引用类型 slice map chan(多个协程沟通的通道 并发使用)
- 接口类型 interface
- 函数类型 func

## 类型零值
- 零值并不等于空值 而是当变量被声明为某种类型后的默认值 通常情况下值类型默认为0 bool 默认为false string默认为空字符串

## 类型转换
- go不存在隐式类型转换 所有的类型转换必须显示声明
- 转换只能发生在两种相互兼容的类型之间
- var a float32 = 1.1; b:= int(a)
- 严格意义上讲type newInt int 这里并不能说是int的别名 而只是底层数据结构相同 在这里称为自定义类型
- 在进行类型转换的时候仍然需要显式转换 但是byte 和 rune 确确实实为unit8 和 int32的别名 可以进行相互转换


 


