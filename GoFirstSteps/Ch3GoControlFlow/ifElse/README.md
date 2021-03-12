## if/else

Go 无需在条件中使用括号, else 为可选项, 不支持三元 if 语句.

一个示例:

```go
package main

func main() {
    x := 27
    if x % 2 == 0 {
        println(x, "is even")
    }
}
```

### 复合 if 语句

支持 `else if` 进行嵌套, 当然更推荐用 `switch`

```go
package main

import "fmt"

func givemeanumber() int {
	return -1
}

func main() {
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// num 作用域结束, 继续使用则报错 undefined: num
	// fmt.Println("num: ", num)
}

```