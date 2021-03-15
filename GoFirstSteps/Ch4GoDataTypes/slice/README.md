## 切片

上回书说到, 数组是切片的基础. 切片同数组一样, 也表示一系列类型相同的元素.
但切片的大小是动态的.

切片是数组或另一个切片之上的数据结构, 我们将源数组或切片称为基础数组.
通过切片, 可访问整个基础数组, 也可以只访问部分元素.

切片包含三个组件:
1. 指向基础数组第一个可访问元素的指针;
2. 长度
3. 容量

### 声明和初始化切片

```go
func main() {
    months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
    fmt.Println(months)
    fmt.Println("Length:", len(months))
    fmt.Println("Capacity:", cap(months))
}
```

和数组相同, 但不用指定长度. 可使用内置函数 `len` 和 `cap` 获取长度和容量.

### 切片项

Go 支持切片运算符 `s[i:p]`, 即获取数组或切片`s`的第`i`到第`p`个这个区间的元素,
其中不包括 `s[p]`.

例如:
```go
func main() {
    months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
    quarter1 := months[0:3]
    quarter2 := months[3:6]
    quarter3 := months[6:9]
    quarter4 := months[9:12]
    fmt.Println(quarter1, len(quarter1), cap(quarter1))
    fmt.Println(quarter2, len(quarter2), cap(quarter2))
    fmt.Println(quarter3, len(quarter3), cap(quarter3))
    fmt.Println(quarter4, len(quarter4), cap(quarter4))
}
```

需要注意的是四个切片的长度相同但容量不同, 取决于基础数组是否有更多位置可供使用.
比如我们可以基于 `quarter2` 创建 `quarter2Extended := quarter2[:4]`,
其值为 `[April May June July]`.

### 追加项

Go 提供了内置函数 `append(slice, element)`, 便于将元素添加到切片.

```go
func main() {
    var numbers []int
    for i := 0; i < 10; i++ {
        numbers = append(numbers, i)
        fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
    }
}
```

运行后可以看到 numbers 的容量不足以容纳时将会翻倍.

### 删除项

Go 没有内置删除元素的函数, 但可以创建一个不含需要删除元素的新的切片.

```go
func main() {
    letters := []string{"A", "B", "C", "D", "E"}
    remove := 2

    if remove < len(letters) {
        fmt.Println("Before", letters, "Remove ", letters[remove])
        letters = append(letters[:remove], letters[remove+1:]...)
        fmt.Println("After", letters)
    }
}
```

另一种方法是创建切片的副本.

### 创建切片的副本

Go 提供了内置函数 `copy(dst, src []Type)` 用于创建切片的副本. 例如:

```go
slice2 := make([]string, 3)
copy(slice2, letters[1:4])
```

因为切片依赖基础数组, 若更改切片, 基础数组会受影响. 因此需要创建副本.

若不使用副本, 更改切片会更改数组:
```go
func main() {
    letters := []string{"A", "B", "C", "D", "E"}
    fmt.Println("Before", letters)

    slice1 := letters[0:2]
    slice2 := letters[1:4]

    slice1[1] = "Z"

    fmt.Println("After", letters)
    fmt.Println("Slice2", slice2)
}
```

使用副本就不受影响:
```go
func main() {
    letters := []string{"A", "B", "C", "D", "E"}
    fmt.Println("Before", letters)

    slice1 := letters[0:2]

    slice2 := make([]string, 3)
    copy(slice2, letters[1:4])

    slice1[1] = "Z"

    fmt.Println("After", letters)
    fmt.Println("Slice2", slice2)
}
```
