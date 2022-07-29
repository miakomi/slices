# Slices tools 

Slices tools is a module written in Golang that facilitates working with slices in Golang.

## Features:
•Reverse slices

•to be continue...


## Quick start:

### Install Slices tools
```
go mod init * 
go get -u github.com/miakomi/slices
go mod tidy
```

### Reverse you slice 
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
