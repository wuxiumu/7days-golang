1. 指针
   区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。
```go
package main

import "fmt"

type Person struct {
    name string
}

// 指针类型的Person
func chageName(p *Person) {
   p.name = "Tom"
}

func main() {
   // 
    P := Person{
 
      name:"Alice"
    }
    // 指针类型的值传递 取址
    chageName(&P)
    fmt.Println(P.name) // 输出修改后的值
}
```
要搞明白Go语言中的指针需要先知道3个概念：指针地址、指针类型和指针取值

Go语言中的函数传参都是值拷贝，而不是引用传递。

当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无须拷贝数据。

要记住两个符号：&（取地址）和*（根据地址取值）。
```go
package main

import "fmt"

func main() {
    var a int = 10
    var ptr *int = &a // 取地址
    fmt.Println(*ptr) // 取值

    *ptr = 20 // 修改值
    fmt.Println(a) // 输出修改后的值
}
```


