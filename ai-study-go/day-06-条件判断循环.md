# if
```go
age := 20
if age >= 18{
    fmt.Println("你已满18周岁")
}else if age >= 16{    
	fmt.Println("你已满16周岁")
}else{
    fmt.Println("你还未满18周岁")
}	
```

# for
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

var i int
for {
	time.Sleep(10)
	fmt.Println(i)
	i++
}
for j<10 {
	fmt.Println(j)
	j++
}

// 99乘法表
for i := 1; i <= 9; i++ {    
    for j := 1; j <= i; j++ {
        fmt.Printf("%d * %d = %d\t", i, j, i*j)
    }
    fmt.Println()
}

for循环 还有一种用法 for range，可以遍历数组、切片、map、字符串等。

name:= "hello world"
nameRune := []rune(name)
for key, value := range name {
    fmt.Println(key, value)
	fmt.Printf("%c\r\n", value)
}

for _, value := range []int{1, 2, 3, 4, 5} {
    fmt.Println(value)
}

// break、continue、goto语句    
round := 0
for {
    if round == 5 {
        break
    }
    fmt.Println("round:", round)
    round++
    if round == 3 {
        continue
    }
    fmt.Println("continue")
    goto L1
}
```

# switch
```go
age := 20
switch {
case age >= 18:
    fmt.Println("你已满18周岁")
case age >= 16:
    fmt.Println("你已满16周岁")
default:
    fmt.Println("你还未满18周岁")
}
``` 

# 注意事项
- 条件判断语句和循环语句都可以嵌套使用。
- switch语句的case后面可以有多个条件，用逗号隔开。
- switch语句的default语句是可选的，如果没有匹配的case语句，则执行default语句。
