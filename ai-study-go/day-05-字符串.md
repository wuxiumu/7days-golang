# 字符串

## 1.转义符
\n 换行符 ASCII 10

\r 回车符 ASCII 13

## 2.输出格式化
```azure
fmt.Printf("Hello, %s!\n", "world")
```

## 3.字符串的拼接
```azure
fmt.Sprintf("Hello, %s!", "world")
fmt.Sprint

var builder strings.Builder
builder.WriteString("Hello, ")
builder.WriteString("world")
re := builder.String()
fmt.Println(re)
```
## 4.字符串的比较
```azure
strings.EqualFold("hello", "HELLO") // true 忽略大小写
```

## 5.字符串的操作
```azure
strings.Contains("hello world", "world") // true 包含子串
strings.Count("hello world", "l") // 3 出现次数
strings.Split("hello world", " ") // ["hello", "world"] 分割字符串
// 字符串是否包含前缀和后缀
strings.HasPrefix("hello world", "he") // true
strings.HasSuffix("hello world", "ld") // true
    

strings.Index("hello world", "l") // 2 子串第一次出现的位置
strings.LastIndex("hello world", "l") // 9 子串最后一次出现的位置  
strings.Join([]string{"hello", "world"}, " ") // "hello world" 连接字符串
strings.Repeat("hello", 3) // "hellohellohello"  重复字符串
strings.Replace("hello world", "l", "L", 2) // "heLLo world" 替换字符串
strings.ToLower("HELLO WORLD") // "hello world" 转小写

``` 