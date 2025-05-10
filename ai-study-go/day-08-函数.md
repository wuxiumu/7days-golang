# 函数
1.变量

2.闭包

3.接口

```go
package main
func add(x, y int) int {
	return x + y,nil
}
func main() {
	sun,_:=add(1,2)
	fmt.Println(sun)
}
// 值传递
```
## 函数的可变参数
```go
func add(x...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func main() {
	fmt.Println(add(1, 2, 3, 4, 5))
}
```

## go函数的闭包
```go   
package main

import "fmt"     // 引入fmt包

func adder() func(int) int { // 定义一个函数作为返回值，返回值为一个函数
	sum := 0
	return func(x int) int { // 定义一个匿名函数，接收一个int参数，返回一个int参数
		sum += x
		return sum
	}
}

func main() {
	f := adder() // 调用adder函数，返回一个匿名函数
	fmt.Println(f(1)) // 调用匿名函数，传入参数1，返回1
	fmt.Println(f(2)) // 调用匿名函数，传入参数2，返回3
	fmt.Println(f(3)) // 调用匿名函数，传入参数3，返回6
}
```

## defer 语句的使用
```go
package main

import "fmt"

func main() {
	defer fmt.Println("world") // 打印"world"后，再执行main函数
	fmt.Println("hello")
}
```

## error的设计理念
1. error接口 
2. painc函数
3. recover函数
```go
if err != nil {
	fmt.Println(err)
} else {
	// 处理正常逻辑 
}	
```


```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)           // 调用divide函数，传入参数10,0
	if err != nil {
		fmt.Println(err)                   // 打印错误信息
	} else {
		fmt.Println(result)                // 打印结果
	}
}   
```
## 计算器
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Enter an expression: ")
	var input string
	fmt.Scanln(&input)
	result, err := calculate(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func calculate(input string) (float64, error) {
	var stack []float64
	var current float64
	var operator string
	for _, c := range input {
		if c == '+' || c == '-' || c == '*' || c == '/' || c == '^' {
			if operator == "" {
				operator = string(c)
			} else {
				result, err := calculateStack(stack, operator)
				if err != nil {
					return 0, err
				}
				stack = []float64{result}
				operator = string(c)
			}
		} else if c == 'x' || c == 'X' {
			operator = "*"
		} else if c == '÷' || c == ':' {
			operator = "/"
		} else if c == '(' {
			stack = append(stack, current)
			current = 0
		} else if c == ')' {
			result, err := calculateStack(stack, operator)
			if err != nil {
				return 0, err
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, result)
		} else {
			num, err := strconv.ParseFloat(string(c), 64)
			if err != nil {
				return 0, err
			}
			current = current*10 + num
		}
	}
	if operator != "" {
		result, err := calculateStack(stack, operator)
		if err != nil {
			return 0, err
		}
		stack = []float64{result}
	}
	if len(stack)!= 1 {
		return 0, errors.New("invalid expression")
	}
	return stack[0], nil
}

func calculateStack(stack []float64, operator string) (float64, error) {
	if len(stack) < 2 {
		return 0, errors.New("invalid expression")
	}
	var result float64
	if operator == "+" {
		result = stack[len(stack)-2] + stack[len(stack)-1]
	} else if operator == "-" {
		result = stack[len(stack)-2] - stack[len(stack)-1]
	} else if operator == "*" {
		result = stack[len(stack)-2] * stack[len(stack)-1]
	} else if operator == "/" {
		if stack[len(stack)-1] == 0 {
			return 0, errors.New("division by zero")
		}
		result = stack[len(stack)-2] / stack[len(stack)-1]
	} else if operator == "^" {
		result = stack[len(stack)-2] ** stack[len(stack)-1]
	}
	stack = stack[:len(stack)-2]
	stack = append(stack, result)
	return result, nil
}
```