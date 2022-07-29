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
slice := []int{1,2,3}
slice2 := slices.Reverse(slice)
fmt.Println(slice2) //output [3 2 1]
```
or without return slice
```go 
slice := []int{1, 2, 3}
slices.ReverseSlice(slice)
fmt.Println(slice) //output [3 2 1]
```
