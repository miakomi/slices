# Slices tools 

Slices tools is a module written in Golang that facilitates working with slices in Golang.

## Features:
•Search a numbers in slice (support int and float32)

•Reverse slices

•to be continue...


## Quick start:

### Install Slices tools
```
go mod init * 
go get -u github.com/miakomi/slices
go mod tidy
```

### Search the index of a numbers in a []int slice
_use only the sorted slice (sort.Ints(slice) on package slice)_
```go
slice := []int{1, 3, 8, 16, 56}
res := slices.SearchInt(slice, 56) //output Ok, i(index) = 4
res = slices.SearchInt(slice, 58) // output NO
```
#### of float32
```go
slice := []float32{1.1, 3.5, 8.0, 16.5, 56.8}
res := slices.SearchFloat32(slice, 56.8) //output Ok, i(index) = 4
res = slices.SearchFloat32(slice, 56.7) // output NO
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
