### 记录框架

例如 `logrus`, `zerolog`, `zap`, `Apex`, 
以下以 `zerolog` 为例.

```go
package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    log.Print("Hey, I'm a log message !")
}
```

日志输出变为了 JSON 格式.

另外可以快速添加上下文数据:

```go
package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    log.Debug().Int("EmployeeID", 1001).Msg("Getting employee information")
    log.Debug().Str("Name", "John").Send()
}
```

`zerolog` 还可实现分级的日志记录, 格式化的堆栈跟踪, 使用多个记录器实例来管理不同输出等.
详细参见 [Github](https://github.com/rs/zerolog)

