1.bool
```azure
var b bool = true
```
2.int
    1.int8 (-128-127)
    2.int16 (-32768-32767)
    3.int32
    4.int64 
    5.uint8 (0-255)
    6.uint16
    7.uint32
    8.uint64
    9.uintptr   
```azure
var i int = 10
var a int8 = -128
var b int16 = -32768
var c int32 = 2147483647
var d int64 = -9223372036854775808
var e uint8 = 255
var f uint16 = 65535
var g uint32 = 4294967295
var h uint64 = 18446744073709551615
var j uintptr = 10

var e int // 动态类型
a = int8(b) // 强制类型转换   

int 类型转换规则：
int32 -> int64 -> uint -> uintptr

```

3.float
    1.float32
    2.float64
```azure
var f1 float32 = 3.14
var f2 float64 = 3.141592653589793

```
4. byte
```azure
纯英文字符的 ASCII 码
type byte uint8 // byte 类型是 uint8 的别名 存放字符的 ASCII 码
0开始
var c byte
c = 'a' // 97
c = '中' // 20013
fmt.Println("c=%c", c) // c=中zw
```
5.rune
```azure
中文字符的 Unicode 码
type rune int32 // rune 类型是 int32 的别名 存放 Unicode 码

```
5.string
6.complex64
7.complex128