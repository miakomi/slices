# Slices tools 

Slices tools is a module written in Golang that facilitates working with slices in Golang.

## Features:
•Search a numbers in a slice (support int/uint and float32/64)

•Reverse slices

•to be continue...


## Quick start:

### Install Slices tools
```
go mod init * 
go get -u github.com/miakomi/slices
go mod tidy
```

### Search the index of a numbers in a slice
_Recommended to use sorted slices because the search function in unsorted slices is slowly_
```go
slice := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
ok, i := slices.Search(slice, 4.4) //return true, 3(index of a number)
//ok, i := slice.Search(slice. 15.4) //return false, -1
if ok {
	fmt.Println(slice[i]) //4.4
}
```

### Reverse slice 
```go 
slice := []int{1, 2, 3}
slices.ReverseSlice(slice)
fmt.Println(slice) //output [3 2 1]
```

#### or float slice.. 
```go 
arr := []float32{1.1, 2.2, 3.3}
slices.Reverse(arr)
fmt.Println(arr) //output [3.3, 2.2, 1.1]
```
#### or string slice.. 
```go 
arr := []string{"one", "two", "three"}
slices.Reverse(arr)
fmt.Println(arr) //output ["tree", "two", "one"]
```
