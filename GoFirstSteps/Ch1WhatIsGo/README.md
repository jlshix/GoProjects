# 什么是 GO?

Go 是一种强类型通用语言, 主要用于后端开发. 许多知名项目入 Docker, Kubernetes 也是使用 Go 语言编写.

Go 与 C 有很多相似之处, 继承了 C 语法的许多方面. 它借鉴并改编了其他编程语言, 同时删除了过于复杂的功能.
如面向对象的编程范例的某些部分.

在 macOS 中使用 `brew install go` 安装, 使用 `go version` 验证是否安装成功.

工作区的配置在 GO 1.11 后可不再遵循 `$GOPATH/src` 的模式. 只需配置环境变量 `GO111MODULE=on` 即可.

一个 helloworld 示例:

新建 `helloworld` 文件夹并进入;

新建 `main.go` 并编写:

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world!")
}
```

在终端执行 `go run main.go` 执行此文件;
在终端执行 `go build main.go` 得到一个 `main` 二进制可执行文件.
