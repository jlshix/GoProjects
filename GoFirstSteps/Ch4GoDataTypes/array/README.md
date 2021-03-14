## 数组

Go 中的数组是一种特定类型且长度固定的数据结构. 必须在声明或初始化时确定大小, 且无法调整大小.
因此数组并不常用, 但它们是切片和映射的基础.

### 声明数组

例如:

```go
func main() {
    var a [3]int
    a[1] = 10
    println(a[0])
    println(a[1])
    println(a[len(a)-1])
}
```

声明后会使用数据类型进行初始化.

`len` 是内置函数, 可以获取数组, 切片和映射的长度.

### 初始化数组

例如:

```go
func main() {
    cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
    fmt.Println("Cities:", cities)
}
```

使用值进行初始化, `cities[4]` 为空字符串.


### 数组中的省略号

如果不知道长度但知道有多少数据, 可以使用省略号作为长度.

```go
func main() {
    cities := [...]string{"New York", "Paris", "Berlin", "Madrid"}
    fmt.Println("Cities:", cities)
}
```

则 `cities` 长度为 4.

另外一种使用方式如下:

```go
func main() {
    numbers := [...]int{99: -1}
    fmt.Println("First Position:", numbers[0])
    fmt.Println("Last Position:", numbers[99])
    fmt.Println("Length:", len(numbers))
}
```

此时数组长度为 100, 因为你为第 99 个位置指定了值, 其余都为默认值 0.

### 多维数组

Go 支持多维数组, 例:

```go
func main() {
    var twoD [3][5]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 5; j++ {
            twoD[i][j] = (i + 1) * (j + 1)
        }
        fmt.Println("Row", i, twoD[i])
    }
    fmt.Println("\nAll at once:", twoD)
}
```

同理也支持三维, 四维...
