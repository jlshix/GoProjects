## 映射

Go 中的映射是一个哈希表, 是一个键值对的集合.
映射中所有的键都必须具有相同的类型, 值也是如此.

### 声明和初始化映射

使用 `map` 关键字定义映射. 格式为 `map[T]T`

```go
func main() {
    studentsAge := map[string]int{
        "john": 32,
        "bob":  31,
    }
    fmt.Println(studentsAge)
}
```

也可使用内置函数 `make` 创建空映射:
`studentsAge := make(map[string]int)`

### 添加项

例如:
```go
func main() {
    studentsAge := make(map[string]int)
    studentsAge["john"] = 32
    studentsAge["bob"] = 31
    fmt.Println(studentsAge)
}
```

若使用 `var studentsAge map[string]int` 则会创建一个 `nil映射`,
无法向其添加元素, 请使用 `make`.

### 访问项

使用下标进行访问, 当下标不存在时会返回数据类型对应的默认值.

```go
func main() {
    studentsAge := make(map[string]int)
    studentsAge["john"] = 32
    studentsAge["bob"] = 31
    fmt.Println("Christy's age is", studentsAge["christy"])
}
```

也可以对存在性进行检查:

```go
func main() {
    studentsAge := make(map[string]int)
    studentsAge["john"] = 32
    studentsAge["bob"] = 31

    age, exist := studentsAge["christy"]
    if exist {
        fmt.Println("Christy's age is", age)
    } else {
        fmt.Println("Christy's age couldn't be found")
    }
}
```

### 删除项

若要删除项, 请使用 Go 内置函数 `delete`.

```go
func main() {
    studentsAge := make(map[string]int)
    studentsAge["john"] = 32
    studentsAge["bob"] = 31
    delete(studentsAge, "john")
    delete(studentsAge, "christy")
    fmt.Println(studentsAge)
}
```

若删除不存在的项, go 不会进入 panic.

### 映射中的循环

可使用基于范围的循环, 例如:

```go
func main() {
    studentsAge := make(map[string]int)
    studentsAge["john"] = 32
    studentsAge["bob"] = 31
    for name, age := range studentsAge {
        fmt.Printf("%s\t%d\n", name, age)
    }
}
```

当然我们也可使用 `_` 变量忽略其中任何一个.

```go
for _, age := range studentsAge {
    ...
}
```

若只需要键则可:

```go
for name := range studentsAge {
    ...
}
```