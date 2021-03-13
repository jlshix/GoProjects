## switch

### 基本语法

像 `if` 语句一样, `switch` 条件不需要括号.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    sec := time.Now().Unix()
    rand.Seed(sec)
    i := rand.Int31n(10)

    switch i {
    case 0:
        fmt.Print("zero...")
    case 1:
        fmt.Print("one...")
    case 2:
        fmt.Print("two...")
    default:
		fmt.Println("no match: ", i)
    }

    fmt.Println("ok")
}
```

### 使用多个表达式

多个匹配操作相同时, 可在 case 后跟多个表达式, 使用逗号分隔.

```go
switch i {
case 0, 2, 4, 6, 8:
    fmt.Println(i, "is even")
case 1, 3, 5, 7, 9:
    fmt.Println(i, "is odd")
}
```

### 调用函数

switch 后可跟函数调用, 以返回值编写 case 语句.

```go
switch time.Now().Weekday().String() {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("It's time to learn some Go.")
default:
    fmt.Println("It's weekend, time to rest!")
}
```

case 中也可以调用函数, 如:

```go
var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

contact := "foo@bar.com"

switch {
case email.MatchString(contact):
    fmt.Println(contact, "is an email")
case phone.MatchString(contact):
    fmt.Println(contact, "is a phone number")
default:
    fmt.Println(contact, "is not recognized")
}
```

此时 switch 后无其他表达式.

### 省略条件

switch 省略条件时要求 case 后表达式为布尔值, 为真则执行:

```go
r := rand.Float64()
switch {
case r > 0.1:
    fmt.Println("Common case, 90% of the time")
default:
    fmt.Println("10% of the time")
}
```

### 使逻辑进入到下一个 case

类 C 的编程语言通常在 case 后需显式使用 `break` 退出 switch.
Go 则相反, 默认执行完匹配到的 case 就退出.

若要使逻辑忽略下个紧邻 case 的判定条件而直接执行, 可使用 `fallthrough` 关键字.

```go
switch num := 15; {
case num < 50:
    fmt.Printf("%d is less than 50\n", num)
    fallthrough
case num > 100:
    fmt.Printf("%d is greater than 100\n", num)
    fallthrough
case num < 200:
    fmt.Printf("%d is less than 200", num)
}
```

此时三个 Printf 都会打印.

