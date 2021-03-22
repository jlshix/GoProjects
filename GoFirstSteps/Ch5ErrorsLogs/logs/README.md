## 日志

日志在排查程序问题时很有用. 最终用户只需要一条消息, 指示程序出现了问题, 而开发人员需要更多信息.

因为我们需要复现并修复问题.

### log 包

log 包作为标准包, 可以像 fmt 一样使用, 不提供日志级别, 且不允许为每个包配置单独的记录器.
若要实现这些功能则需要记录框架.

```go
import "fmt"

func main() {
    log.Print("Hey, I'm a log!")
}
```

运行则看到一行带有日期前缀的输出.

因为默认情况下 `log.Print` 将日期和时间添加为日志消息的前缀.

可以使用 `log.Fatal` 记录错误并退出程序, 如图 `os.Exit(1)` 一样.

```go
func main() {
    log.Fatal("Hey, I'm an error log !")
    fmt.Print("Can you see me ?")
}
```

后一行未运行, 因为前一行就已经停止了程序.

若将上例中的 `log.Fatal` 替换为 `log.Panic` 则会调用 `panic`, 给出错误堆栈跟踪.

使用 `log.SetPrefix` 则可向程序的日志消息添加前缀.

### 记录到文件

日志除了打印在控制台之外还可以记录到文件, 以便稍后或实时处理这些日志.

```go
func main() {
    file, err := os.OpenFile("info.log", os.CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)
    log.Print("Hey, I'm a log!")
}
```

运行以上代码则控制台中无任何内容, 日志会存放到文件中.

