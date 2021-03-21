## 如何在 Go 中处理错误

一些无法控制的外部因素导致程序运行失败如果可以预见, 那就能再发生时编写代码解决问题.

Go 的异常处理与其他语言不同, 错误处理也是如此. Go 中可能失败的函数应始终返回一个额外值.

Go 使用 panic 和 recover 管理异常或意外行为. 但错误是已知失败, 程序应该可以处理.

Go 的错误处理方法使用 if 和 return 控制, 如:

```go
employee, err := getInformation(1000)
if err != nil {
    // something is wrong, do something
}
```

如上, 当函数返回错误时, 错误通常为最后一个值. 调用方负责检查是否存在错误并处理错误.

```go
func getInformation(id int) (*Employee, error) {
    employee, err := apiCallEmployee(id)
    if err != nil {
        return nil, fmt.Errorf("Got an error: %v", err)
    }
    return employee, nil
}
```

若错误为暂时性错误, 可以选择重试:

```go
func getInformation(id int) (*Employee, error) {
    for tries := 0; tries < 3; tries ++ {
        employee, err := apiCallEmployee(id)
        if err == nil {
            return employee, nil
        }
        fmt.Println("Server is not responding, retrying ...")
        time.Sleep(time.Second * 2)
    }
    return nil, fmt.Errorf("server has failed to respond to get the employee information")
}
```

### 创建可重用的错误

使用 `errors.New()` 函数创建错误并重用.

```go
var ErrNotFound = errors.New("Employee not found")

func getInformation(id int) (*Employee, error) {
    if id != 1001 {
        return nil, ErrNotFound
    }
    
    employee := Employee{LastName: "Snow", FirstName: "John"}
    return &Employee, nil
}
```

按照惯例, 错误变量使用 `Err` 前缀.

使用 `errors.Is()` 函数确认获得错误的类型, 如:

```go
employee, err := getInformation(1000)
if errors.Is(err, ErrNotFound) {
    fmt.Println("NOT FOUND %v\n", err)
} else {
    fmt.Println(employee)
}
```

### 推荐做法

1. 始终检查是否存在错误, 以免向用户公开不必要的信息

2. 在错误消息中包含前缀来了解错误的来源

3. 创建尽可能多的可重用错误变量

4. 了解使用错误和 panic 之间的差异. 不能执行其他操作时再使用 panic

5. 在记录错误时记录尽可能多的详细信息, 并打印出最终用户能够理解的错误