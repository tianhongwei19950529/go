# Go 命令行工具
`go xxx`

# 代码组织
go 程序被组织成包

# 模块
一系列相关的Go包，放到一个目录里，就组成了一个模块

github go项目仓库中，**通常只包含一个模块**，但其实多个也可以的。

**在每个模块的根目录中要放置他的 go.mod 文件**

```
- module
    - package0
        - file1.go
        - file2.go
    - package1
        - ...
    - package2
        - ...
    - go.mod
```

# go.mod 内容

```
module github.com/xxx/asd      // 第一行，该模块的模块路径

go 1.5          // 第二行，go版本

require (
    
)
```

# demo
创建模块
```sh
cd hello
go mod init example/user/hello
```

然后会产生 go.mod 文件。接下来创建 hello.go 里写 Hello World（注意 package main ）不要丢
```
    - hello
        - hello.go
        - go.mod
```

进行编译安装，`go install example/user/hello`，将模块编译成叫 hello 的二进制文件 并安装在了家目录的go/bin/里 `ll ~/go/bin/hello`

你可以通过更改 GOPATH 和 GOBIN 两个环境变量来自定义安装的位置。如果设置GOPATH会安装到GOPATH列表中第一个目录的bin子目录，如果设置GOBIN则会直接安装到GOBIN中

设置值：`go env -w  GOBIN=/somewhere`   恢复出厂设置：`go env -u`


go install 注意地方：如果你当前工作目录不在 example/user/hello 模块里面，go install 命令会执行失败。如果你工作目录在模块中，`go install example/user/hello`, `go install .`, `go install` 三条命令等价

继续在 hello 模块里添加 morestrings 包(创建此文件夹)，创建 reverse.go 文件，在里面编写函数实现字符串反转操作。
reverse.go 内容：
```go
// Package morestrings xxxx
package morestrings


// ReverseRunes retures xxx
func ReverseRunes(s string) string {
	r := []rune(s)
	for i,j := 0,len(r)-1; i < len(r)/2; i,j = i+1,j-1 {
        r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

然后 `go build`，什么也没有发生？不，只有 main 包的 main 函数可以编译出可执行文件。严格来说，这个包还是被编译的了，编译结果被放在了本地的编译缓存中，方便以后取用，而且 `go build` 没报错说明包中没有编译的错误，可以通过这种办法简单验证自己写的代码有没有语法问题。

接下来可在 main 包(hello.go)中调用这个新写的东西了。
```go
package main

import (
	"example/user/hello/morestrings"
	"fmt"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello World"))
}
```


# 远程包
依赖包在 go.mod 文件的 require 中，使用方法直接 import 即可。

`go mod tidy` 该命令会自动安装满足依赖。依赖的远程包会自动下载到 `$GOPATH/pkg/`，全是只读文件

移除下载：`go clean -modcache`

镜像代理：`go env -w GO111MODULE=on`  `go env -w GOPROXY=https://goproxy.cn,direct`，设置完成之后最好在重启一下终端，如果出警告设置之前先执行`unset GOPROXY`


# 测试
写 `xxx_test.go` 文件，导入 testing，测试函数 TestXxx(t *testing.T)，设置好一些输入输出用例，没通过就`t.Error()`或`t.Fail()`

reverse_test.go
```go
package morestrings

import "testing"

func TestReverseRunes(t *testing.T)  {
	cases := []struct{
		in, want string
    } {
		{"aaa", "aaa"},
		{"abc", "cba"},
    }
	for _, c := range cases {
        got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
        }
	}
}
```