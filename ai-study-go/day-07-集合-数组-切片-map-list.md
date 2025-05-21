```
slice1 := []int{1, 2, 3, 4, 5}
// 数组 长度固定，性能比较高
arr := [5]int{1, 2, 3, 4, 5}
course := [...]string{"Go", "Python", "Java"}
// 多维数组
var courseInfo [3][4]string
courseInfo[0][0] = "Go语言"
courseInfo[0][1] = "2021-10-10"
courseInfo[0][2] = "19:00"
courseInfo[0][3] = "40"
courseInfo[1] = [4]string{"Python语言", "2021-10-11", "19:30", "30"}
// 切片 长度可变，性能比较高
slice2 := slice1[1:4]
// 切片的长度
len(slice1)
// 切片的容量
cap(slice1)
// 切片的元素类型
slice1[0]
// 切片的迭代
for i := 0; i < len(slice1); i++ {
    fmt.Println(slice1[i])
}
// 切片的切片
slice3 := slice1[1:4]     // [2 3 4]    
slice4 := slice3[1:3]     // [3]        
slice5 := slice4[0:1]     // [3]        
slice6 := append(slice5, 6) // [3 6]       
// 切片的拼接    
slice7 := append(slice1, slice2...) // [1 2 3 4 5 2 3 4]    
// 数组的拼接    
arr2 := [3]int{6, 7, 8}    
arr3 := [5]int{0, 0, 0, 0, 0}    
arr3 = append(arr3[:2], arr2[:]...)    
// 切片的复制    
slice8 := make([]int, len(slice1))    
copy(slice8, slice1)    
// 数组的复制    
arr4 := [5]int{1, 2, 3, 4, 5}    
arr5 := arr4    
// 切片的比较    
slice1 == slice2 // false    
// 数组的比较    
arr == arr2 // false    
// 切片的排序    
sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })    
// 数组的排序    
sort.Sort(sort.IntSlice(arr))    
// 切片的查找    
index := sort.Search(len(slice1), func(i int) bool { return slice1[i] == 3 })    
// 数组的查找    
index := sort.Search(len(arr), func(i int) bool { return arr[i] == 3 })    
// 切片的删除    
slice9 := append(slice1[:2], slice1[3:]...)    
// 数组的删除    
arr6 := [5]int{1, 2, 4, 5, 6}    
arr7 := arr6    
// 切片的元素更新    
slice1[0] = 10    
// 数组的元素更新    
arr[0] = 10    
// 切片的元素追加    
slice1 = append(slice1, 11)    
// 数组的元素追加    
arr = append(arr, 11)    
// 切片的元素删除    
slice1 = append(slice1[:2], slice1[3:]...)    
// 数组的元素删除    
arr = append(arr[:2], arr[3:]...)    
// 切片的元素插入    
slice1 = append(slice1[:2], 10, slice1[2:]...)    
// 数组的元素插入    
arr = append(arr[:2], 10, arr[2:]...)    
// 切片的元素替换    
slice1 = append(slice1[:2], 10, slice1[3:]...)    
// 数组的元素替换    
arr = append(arr[:2], 10, arr[3:]...)    
// 切片的元素查找    
index := sort.Search(len(slice1), func(i int) bool { return slice1[i] == 3 })    
// 数组的元素查找    
index := sort.Search(len(arr), func(i int) bool { return arr[i] == 3 })    
// 切片的元素删除    
slice1 = append(slice1[:index], slice1[index+1:]...)    
// 数组的元素删除    
arr = append(arr[:index], arr[index+1:]...)    
// 切片的元素插入    
slice1 = append(slice1[:index], 10, slice1[index:]...)    
// 数组的元素插入    
arr = append(arr[:index], 10, arr[index:]...)    
// 切片的元素替换    
slice1 = append(slice1[:index], 10, slice1[index+1:]...)    
// 数组的元素替换    
arr = append(arr[:index], 10, arr[index+1:]...)     
// make函数创建切片    
slice1 := make([]int, 5)    
// make函数创建数组    
arr := make([5]int, 5)    
// 复制切片 
copy(slice1, slice2)    
// 复制数组 
arr2 := arr    
## 面试问题summary
1. 请解释一下切片的三个属性：长度、容量、元素类型。
2. 请解释一下数组和切片的区别。
3. 请解释一下make函数的作用。
4. 请解释一下切片的拼接、复制、比较、排序、查找、删除、插入、替换、查找、删除、插入、替换的操作。
5. 请解释一下数组的拼接、复制、比较、排序、查找、删除、插入、替换的操作。
```
结构体
1. 请解释一下结构体的定义。 
```go
// 结构体值传递 
type Person struct {
    Name string
    Age int
}

```
2. 请解释一下结构体的字段。
3. 请解释一下结构体的指针。
4. 请解释一下结构体的嵌套。
```go
package main

import (
    "fmt"
    "sort"
)
func printSlice(data []string) {
	
}
func main() {   

```