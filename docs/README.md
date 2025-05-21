# golang 基础

本目录由ai生成

现在时间：2025-05-21 16:32:22

1. [go实现一个web文件管理器](./demo/go-web-file-manager/)

## 1. 变量
```
变量定义：
1. 声明变量：var 变量名 变量类型 = 变量初始值
2. 声明多个变量：var (
   变量名1 变量类型1 = 变量初始值1
   变量名2 变量类型2 = 变量初始值2
   ...
   )
3. 声明变量并赋值：var 变量名 变量类型 = 变量初始值
4. 声明变量并省略类型：var 变量名 = 变量初始值
5. 声明变量并省略类型和初始值：var 变量名

## 2. 常量
1. 定义常量：const 常量名 = 常量值
2. 定义多个常量：const (
   常量名1 = 常量值1
   常量名2 = 常量值2
   ...
   )
3. 定义常量并省略类型：const 常量名 = 常量值
4. 定义常量并省略类型和初始值：const 常量名
```

## 3. 类型
```
1. 基本类型：bool、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、float32、float64、complex128、byte、rune、string
2. 复合类型：数组、结构体、指针、函数、切片、接口、map、channel
3. 类型别名：type 类型别名 = 原类型
4. 类型断言：x.(类型)
5. 类型转换：T(x)
6. 零值：bool、数值类型、字符串、指针、接口、函数、切片、map、channel
7. 类型判断：x, ok := x.(类型)
8. 类型选择：switch x.(type) {
   case 类型1:
      // 代码块1
   case 类型2:
      // 代码块2
   default:
      // 代码块3
   }
9. 内建函数：make、len、cap、new、append、copy、close、delete、panic、recover
10. 常见的包：fmt、strconv、time、math、net、os、io、bufio、path、regexp、sort、sync、atomic、runtime、syscall、database/sql、encoding/json、encoding/xml、encoding/binary、encoding/csv、encoding/gob、encoding/xml、crypto/sha1、crypto/md5、crypto/rand、crypto/cipher、crypto/des、crypto/aes、crypto/rsa、crypto/dsa、crypto/ecdsa、crypto/x509、crypto/tls、net/http、net/smtp、net/url、net/rpc、net/rpc/jsonrpc、net/http/httputil、net/http/cgi、net/http/cookiejar、net/http/fcgi、net/http/httptest、net/http/httptrace、net/http/httpproxy、net/http/httputil、net/smtp/smtp、net/
```

## 4. 运算符
```
1. 算术运算符：+、-、*、/、%、++、--、&、|、^、<<、>>、&^
2. 关系运算符：==、!=、<、<=、>、>=
3. 逻辑运算符：&&、||、!
4. 赋值运算符：=、+=、-=、*=、/=、%=、<<=、>>=、&=、|=、^=
5. 条件运算符：x ?  y  :  z
6. 位运算符：&、|、^、~、<<、>>、&^
7. 其他运算符：:=、()、[]、.、->、<-、++、--、,、:=、...
8. 运算符优先级：()、[]、.、->、++、--、&^、^、|、^、&、<<、>>、==、!=、<、<=、>、>=、&&、||、?:、:=、,、...
9. 运算符短路：&&、||、?:
10. 运算符的溢出：int32、int64、uint32、uint64、float32、float64
```
## 5. 控制结构
```
1. if语句：if 条件表达式 {
   代码块1
   } else if 条件表达式 {
   代码块2
   } else {
   代码块3
   }
2. switch语句：switch 表达式 {
   case 值1:
      代码块1
      break
   case 值2:
      代码块2
      break
   default:
      代码块3
   }
3. for语句：for 表达式1; 表达式2; 表达式3 {
   代码块
   }
4. for-range语句：for 变量 := range 表达式 {
   代码块
   }
5. select语句：select {
   case 表达式1:
      代码块1
   case 表达式2:
      代码块2
   default:
      代码块3
   }
6. goto语句：goto 标签名
7. break语句：break 语句
8. continue语句：continue 语句
9. defer语句：defer 代码块
10. 标签语句：标签名: 代码块
```
## 6. 函数
```
1. 函数定义：func 函数名(参数列表) (返回值列表) {
   代码块
   }     
2. 函数调用：函数名(参数列表)
3. 匿名函数：func(参数列表) (返回值列表) {
   代码块
   }
4. 闭包：函数嵌套函数
5. 递归函数：函数调用自身
6. 多返回值：func 函数名(参数列表) (返回值1, 返回值2,...) {
   代码块
   }
7. 错误处理：error接口
8. 接口：接口定义、接口实现、接口类型断言、接口赋值、接口方法
9. 反射：反射包
10. 并发编程：Goroutine、Channel、锁、条件变量、WaitGroup、Context、超时控制、并发原语
11. 单元测试：测试用例、测试函数、测试套件、测试框架
12. 性能分析：CPU、内存、GC、PProf、Trace、火焰图
13. 其他：defer、panic、recover、接口、反射、并发、单元测试、性能分析、错误处理、标签、匿名函数、闭包、递归函数、多返回值、超时控制、Context、WaitGroup、Goroutine、Channel、锁、条件变量、测试用例、测试函数、测试套件、测试框架、CPU、内存、GC、PProf、Trace、火焰图
```

## 7. 包

```
1. 包定义：package 包名
2. 导入包：import "包路径"
3. 包变量：var 变量名 变量类型 = 变量初始值
4. 包常量：const 常量名 = 常量值
5. 包函数：func 函数名(参数列表) (返回值列表) {
   代码块
   }
6. 包初始化：func init() {
   代码块
   }
7. 包测试：func TestXXX(t *testing.T) {
   代码块
   }
```

## 8. 错误处理

```
1. 错误定义：type 错误名 struct {
   错误信息 string
   其他字段  ...
   }
2. 错误检查：if err!= nil {
   代码块
   }
3. 错误打印：fmt.Println(err.Error())
4. 错误类型断言：if 变量名, ok := err.(*错误类型); ok {
   代码块
   }
5. 错误赋值：err = errors.New("错误信息")
```
## 9. 并发编程

```
1. Goroutine：go 函数名(参数列表)
2. Channel：make(chan 类型, 容量)
3. 锁：sync.Mutex、sync.RWMutex、sync.Once  
4. 条件变量：sync.Cond
5. WaitGroup：sync.WaitGroup
6. Context：context.WithCancel、context.WithTimeout、context.Background
7. 超时控制：time.After、time.Sleep
8. 并发原语：atomic.AddInt32、atomic.CompareAndSwapInt32、sync.WaitGroup、sync.Mutex、sync.RWMutex、sync.Once、sync.Cond、context.WithCancel、context.WithTimeout、context.Background、time.After、time.Sleep
9. 其他：Goroutine、Channel、锁、条件变量、WaitGroup、Context、超时控制、并发原语
```
## 10. 反射

```
1. 反射包：reflect.TypeOf、reflect.ValueOf、reflect.Call、reflect.New、reflect.MakeSlice、reflect.MakeMap、reflect.Indirect、reflect.ValueOf、reflect.Value.Interface、reflect.Value.Elem、reflect.Value.FieldByName、reflect.Value.FieldByIndex、reflect.Value.SetInt、reflect.Value.SetFloat、reflect.Value.SetString、reflect.Value.SetBool、reflect.Value.SetBytes、reflect.Value.SetPointer、reflect.Value.SetMapIndex、reflect.Value.SetSliceIndex、reflect.Value.Set
2. 其他：反射包、接口、并发、单元测试、性能分析、错误处理、标签、匿名函数、闭包、递归函数、多返回值、超时控制、Context、WaitGroup、Goroutine、Channel、锁、条件变量、测试用例、测试函数、测试套件、测试框架、CPU、内存、GC、PProf、Trace、火焰图
```
## 11. 单元测试

```
1. 测试用例：func TestXXX(t *testing.T) {
   代码块
   }
2. 测试函数：func 函数名(t *testing.T) {
   代码块
   }
3. 测试套件：package 包名
   import "testing"
   func TestXXX(t *testing.T) {
   代码块   
   }
   func TestYYY(t *testing.T) {
   代码块   
   }
4. 测试框架：go test -v -cover -coverprofile=cover.out ./...
5. 性能分析：go tool pprof -http=:8080 cover.out
```
## 12. 性能分析

```
1. CPU：runtime.NumCPU、runtime.GOMAXPROCS、runtime.NumCgoCall、runtime.ReadMemStats、runtime.GC、runtime.SetBlockProfileRate、runtime.SetMutexProfileFraction、runtime.Stack
2. 内存：runtime.MemStats、runtime.ReadMemStats、runtime.GC、runtime.SetGCPercent、runtime.SetBlockProfileRate、runtime.SetMutexProfileFraction、runtime.Stack
3. GC：runtime.GC、runtime.SetGCPercent、runtime.SetBlockProfileRate、runtime.SetMutexProfileFraction、runtime.Stack
4. PProf：go tool pprof -http=:8080 cover.out
5. Trace：runtime.SetBlockProfileRate、runtime.SetMutexProfileFraction、runtime.Stack
6. 火焰图：go get -u github.com/google/pprof/... && go tool pprof -http=:8080 http://localhost:8080/debug/pprof/profile
7. 其他：CPU、内存、GC、PProf、Trace、火焰图
```
## 13. 其他
```
1. defer：defer 代码块
2. panic：panic("错误信息")
3. recover：recover()
4. 接口：接口定义、接口实现、接口类型断言、接口赋值、接口方法
5. 反射：反射包
6. 并发编程：Goroutine、Channel、锁、条件变量、WaitGroup、Context、超时控制、并发原语
7. 单元测试：测试用例、测试函数、测试套件、测试框架
8. 性能分析：CPU、内存、GC、PProf、Trace、火焰图
9. 错误处理：error接口
10. 标签语句：标签名: 代码块
11. 匿名函数：func(参数列表) (返回值列表) {
   代码块
   }
12. 闭包：函数嵌套函数
13. 递归函数：函数调用自身
14. 多返回值：func 函数名(参数列表) (返回值1, 返回值2, ...) {
   代码块
   }
15. 超时控制：time.After、time.Sleep
16. Context：context.WithCancel、context.WithTimeout、context.Background
17. WaitGroup：sync.WaitGroup
18. 锁：sync.Mutex、sync.RWMutex、sync.Once
19. 条件变量：sync.Cond
```

