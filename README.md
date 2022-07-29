# Arrays tools 

Arrays tools is a module written in Golang that facilitates working with arrays and slices in Golang.

## Features:
•Reverse arrays and slices

•to be continue...


## Quick start:

### Install ArrayTools
```
go mod init * 
go get -u github.com/miakomi/arrays
go mod tidy
```

### Reverse you slice 
```go
slice := []int{1,2,3}
slice2 := arrays.ReverseSlice(slice)
fmt.Println(slice2) //output [3 2 1]
```
or without return slice
```go 
slice := []int{1, 2, 3}
arrays.ReverseSlice(slice)
fmt.Println(slice) //output [3 2 1]
```
