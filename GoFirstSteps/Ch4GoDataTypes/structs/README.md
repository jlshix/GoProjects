## 结构

Go 中的结构可包含零到多个任意类型的字段, 并将它们表示为单个实体.

### 声明和初始化结构

使用 `struct` 关键字.

```go
type Employee struct {
    ID int
    FirstName string
    LastName string
    Address string
}

// 直接声明, 使用类型默认值初始化
var john Employee

// 使用给定值初始化
employee := Employee{1001, "John", "Snow", "Snow Street"}

// 初始化部分字段, 顺序不重要, 未提及的字段使用默认值初始化
employee1 := Employee{LastName: "Snow", FirstName: "John"}

// 字段赋值与访问
employee.ID = 1002
fmt.Println(employee.FirstName)

// 使用 `&` 生成指向结构实例的指针
employeeCopy := &employee
// 修改会影响原来的数据
employeeCopy.FirstName = "David"
fmt.Println(employee)
```

### 结构嵌入

Go 中的结构可以嵌套, 比如定义员工和合同工. 可先定义 `Person` 结构然后重用.

```go
type Person struct {
    ID        int
    FirstName string
    LastName  string
    Address   string
}

type Employee struct {
    Information Person
    ManagerID int
}

var employee Employee
employee.Information.FirstName = "John"
```

但这样会破坏之前的调用代码. 但我们可以这样:

```go
type Employee struct {
    Person
    ManagerID int
}

type Contractor struct {
    Person
    CompanyID int
}

employee := Employee {
    Person: Person{
        FirstName: "John",
    },
}
employee.LastName = "Snow"
fmt.Println(employee.LastName)
```

`Person` 中的所有字段都会嵌入到 `Employee` 中. 因此可以在无需指定 `Person` 字段的情况下
访问 `Employee` 中所有 `Person` 的字段.

但在初始化时必须明确要给哪些字段赋值.


### 使用 JSON 编码和解码结构

Go 对 JSON 有很好的支持, 已包含在官方库中. 还可以执行一些操作, 如重命名结构中字段的名称.
例如:

```go
type Person struct {
    ID        int
    FirstName string `json:"name"`
    LastName  string
    Address   string `json:"address,omitempty"`
}
```

使用 `json.Marshal` 和 `json.Unmashal` 在结构和JSON之间转换.
如:

```go

package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    ID        int
    FirstName string `json:"name"`
    LastName  string
    Address   string `json:"address,omitempty"`
}

type Employee struct {
    Person
    ManagerID int
}

type Contractor struct {
    Person
    CompanyID int
}

func main() {
    employees := []Employee{
        Employee{
            Person: Person{
                LastName: "Snow", FirstName: "John"
            }
        },
        Employee{
            Person: Person{
                LastName: "North", FirstName: "King"
            }
        }
    }

    data, _ := json.Marshal(employees)
    fmt.Println(data)

    var decoded []Employee
    json.Unmarshal(data, &decoded)
    fmt.Println(decoded)
}
```
