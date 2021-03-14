## for

### 基本 for 循环语法

除条件无需括号外, 与 C 语言相同.

分号(;)分隔 for 循环的三个组件:

- 在第一次迭代之前执行的初始语句;
- 在每次迭代之前计算的条件表达式, 该条件为 false 时, 循环会停止;
- 在每次迭代结束时执行的后处理语句.

例如:

```go
func main() {
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("sum of 1..100 is", sum)
}
```

### 空的预处理语句和后处理语句

Go 未提供 `while` 关键字, 当预处理语句和后处理语句留空时, 行为和其他语言中的 while 一致.

```go
func main() {
    var num int64
    rand.Seed(time.Now().Unix())
    for num != 5 {
        num = rand.Int63n(15)
        fmt.Println(num)
    }
}
```

### 无限循环和 break 语句

等同于其他语言的 `while 1 {...}`, 例如:

```go
func main() {
    var num int32
    sec := time.Now().Unix()
    rand.Seed(sec)

    for {
        fmt.Print("Writting inside the loop...")
        if num = rand.Int31n(10); num == 5 {
            fmt.Println("finish!")
            break
        }
        fmt.Println(num)
    }
}
```

### Continue 语句

同其他语言, Go 也提供了行为相同的 `continue`, 用于跳过循环的当前迭代. 例如:

```go
func main() {
    sum := 0
    for num := 1; num <= 100; num++ {
        if num%5 == 0 {
            continue
        }
        sum += num
    }
    fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)
}
```
