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


# 常量的定义
- 常量的值在编译的时候就已经确定
- 常量的定义格式与变量基本相同
- 等号右侧必须是常量或者常量表达式
- 常量表达式中的函数必须是内置函数

# 运算符
- go语言的运算符均是从左到右结合
## 优先级 从高到低
```
^ !   一元运算符
* / % << >> & &^
+ - | ^  二元运算符
== != < <= >
<- 专门用于channel  
（&&） AND
||
```
- & 两个为1 则为1
- | 一个为1 则为1
- ^ 两个只有一个为1 则为1
- &^ A&^B 当B的某一位为1的时候 将A的对应位改为0 结果就是A的最后的值
- && 如果第一表达式为true 则后面的表达式不再执行

 # 指针
 - go虽然保留了指针，但是与其他编程语言不同的是，在go当中不支持指针运算以及->运算服，而直接采用.运算符来操作指针目标对象的成员
 - 操作符&取变量地址 使用*通过指针间接访问目标对象
 - 默认值为 nil  非 NULL
 - ++ -- 作为语句 不做为表达式  只能 a++  这样 b = a++  是不允许的
 
 # 条件表达式 
 ## if
 - if 没有括号  比如 if 1 < 2 {} 
 - if a:= 1; a > 1  a的作用于只在if else中有效  
 - if {  必须同一行
 
 ## 循环 只有 for
 - for {
 
 ## switch
 - 不需要break 一旦符合条件立即终止
 - 如果希望继续执行下一个case  需要使用fallthrough
 - 支持一个初始化的表达式
 - 左大括号必须和switch一行
 
 
## 跳转语句
- goto break continue
- 三个语法都可以配合标签使用
- 标签区分大小写
- break continue 配合使用可以跳出多层循环
- goto 调整执行位置
- 放标签 必须使用 否则报错

# 数组
- 数组指针 是指针 指向数组 
- 指针数组 是数组 里面放的是指针
- go中数组传的值传的是拷贝
- 数组之间可以使用 == 或者 ！= 进行比较  不能 > < 进行比较
- 可以使用new 来创建数组 返回一个数组指针
- go支持多维数组

# 切片 slice  -> C++ vector
- 其本身并不是数组 它指向底层的数组
- 作为变长数组的替代方案 可以关联底层数组的局部或者全部
- 为引用类型
- 可以直接创建或者从底层数组获取生成
- 使用len()获取元素个数 cap()获取容量
- 一般使用make创建
- 如果多个slice指向相同的底层数组 其中一个的值改变会影响全部
- make([]T, len, cap)
- 其中cap可以省略 则和len的值相同
- len表示存数的元素个数 cap表示容量

## Reslice
- 从一个slice中获取一个新的slice
- Reslice时所有以被slice的切片为准
- 索引不能超过被实力测得切片的容量的cap值
- 索引越界不会导致底层数组的重新分配而是引发错误
## Append
- 尾部追加元素
- 可以将一个slice追加在另一个slice的尾部
- 如果最终的长度未超过slice的容量则范湖原始slice，将会重新分配数组并拷贝数据
## Copy
- 相互拷贝的两个slice互不影响

# map
- 类似其他语言中的哈希表或者字典 以key-value的形式存储数组
- key必须是支持==或者!=比较运算的类型 不可以是函数 map 或者 slice
- map的查找比线性搜索快很多 但是比使用索引访问数据的类型慢100倍
- map使用make创建 支持 := 这种简写方式
- map([keyType]valueType, cap) cap表示容量 可省略
- 超出容量的时候会自动扩容 但尽量提供一个合理的初始值
- 使用len()获取元素个数
- 键值对不存在时自动添加 使用delete()删除某键值对
- 使用for range 对map 和 slice进行迭代操作
- 能使用slice或者数组就不要使用map

# function
- go函数不支持嵌套 重载 和 默认参数
- 支持以下特征:
- 无需声明原型 不定长度变参 多返回值 命名返回值参数 匿名函数 闭包
- 定义函数使用关键字func 且左大括号不能另起一行
- 函数也可以作为一种类型使用

## defer
- defer的执行方式类似其他语言中的析构函数 在函数体执行结束后按照调用顺序的相反顺序逐个执行
- 即使函数发生严重错误也会执行
- 支持匿名函数的调用
- 常用于资源清理 文件关闭 解锁以及记录时间等操作
- 通过与匿名函数配合可在return之后修改函数计算结果
- 如果函数体内某个变量作为defer时函数的参数 则在定义defer时即已经获得了值拷贝，
- 如果不是参数，使用时已经在函数体内定义了改参数 则是引用某个变量的地址
- go没有异常机制 但是panic/recover模式来处理错误
- panic可以在任何地方引发 但recover只有在defer调用的函数中有效
  
  


