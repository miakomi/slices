# Slices tools 

Slices tools is a module written in Golang that facilitates working with slices in Golang.

## Features:
•Search in a slice

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
_returns a slice with indexes of search values_
```go
slice := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
ok, i := slices.Search(slice, 4.4) //return true, [3]

if ok {
	fmt.Println(slice[i]) //[3]
}
```
#### or string.. 
```go 
slice := []string{"Alex", "Jon", "Joshua", "Alex"}
ok, i := slices.Search(slice, "Alex") //return true, [0, 3]
  
if ok {
  fmt.Println(i) //[0, 3]
}
```

### Reverse slice 
```go 
slice := []int{1, 2, 3}
slices.Reverse(slice)
fmt.Println(slice) //output [3 2 1]
```

#### or float slice.. 
```go 
slice := []float32{1.1, 2.2, 3.3}
slices.Reverse(slice)
fmt.Println(slice) //output [.3, 2.2, 1.1]
```
#### or string slice.. 
```go 
arr := []string{"one", "two", "three"}
slices.Reverse(slice)
fmt.Println(slice) //output ["tree", "two", "one"]
```
