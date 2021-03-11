## 包

和其他语言一样, 代码可以打包并在其他位置重用.

### main 包

main 包是默认的包, 使用 main 包时会生成二进制文件. 当程序不是 main 包的一部分时,
生成的是包存档文件(拓展名为 .a).

按照约定, 包名与导入路径的最后一个元素相同.

如 `import "fmt"` 则可以使用 `fmt.Println()`;
`import "math/rand"` 则可以使用 `rand.Int()`


### 创建包

我们可以创建一个 `calculator` 文件夹并进入, 再创建一个 `sum.go` 并编写.
第一行为 `package calculator` 指定其包名称.

不同于 Java 使用 `public`, `private` 等关键字标识是否可以从包的内外部调用变量或函数,
Go 遵循以下两个简单规则:
- 小写字母开始的为专用内容;
- 大写字母开始的为公共内容.

示例:

```go
package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func internalSum(number int) int {
    return number - 1
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
    return number1 + number2
}
```

可使用 `go build sum.go` 确认是否正常.


### 创建模块

若要创建模块, 则在 `calculator` 目录下执行以下命令:

`go mod init github.com/myuser/calculator`

运行此命令后, `github.com/myuser/calculator` 就变为了包名, 其他程序中可以使用此名称引用.
同时还会创建一个 `go.mod` 文件, 此时 `calculator` 文件夹正式成为了一个包.


### 引用本地包(模块)

创建 `calculator` 的同级文件夹 `helloworld` 并进入, 编写文件 `main.go` 内容如下:

```go
package main

import "github.com/myuser/calculator"

func main() {
    total := calculator.Sum(3, 5)
    println(total)
    println("Version: ", calculator.Version)
}
```

此时执行 `go run main.go` 会报 `cannot find module providing package github.com/myuser/calculator: working directory is not part of a module`

为此需要我们执行 `go mod init helloworld` 创建一个新的 `go.mod` 文件.

内容为:
```
module helloworld

go 1.15
```

由于我们想要使用 `github.com/myuser/calculator` 的本地副本, 因此我们需要新增两行:

```
require github.com/myuser/calculator v0.0.0

replace github.com/myuser/calculator => ../calculator
```

此时即可正常执行.


### 发布包

大多数使用 github 发布包, 所以我们可以经常看到 import 后包含 "github.com" 的引用.

因此若要将 calculator 发布到 github 账户, 则需要将 `myuser` 改为自己的用户名,
并创建一个名为 `calculator` 的 repo.

包的版本可以通过 git tag 来指定. 这样拉取时就会使用对应的版本.


### 引用外部(第三方)包

只需在其 go.mod 下使用 require 指定, 即可在编译时使用.

如:

```
module helloworld

go 1.15

require (
    github.com/myuser/calculator v0.0.0
    rsc.io/quote v1.5.2
)

replace github.com/myuser/calculator => ../calculator
```