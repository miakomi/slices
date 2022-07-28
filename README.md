# Arrays tools 

Arrays tools is a module written in Golang that facilitates working with arrays and slices in Golang.

## Features:
•Reverse arrays and slices

•to be continue...


## Quick start:

### Install ArrayTools
```
go mod init 
go get -u github.com/Miakomi/arrays
go mod tidy
```

### Reverse you array or slice 
```golang
slice := []int{1,2,3}
slice2 := arrays.Reverse(slice)
fmt.Println(slice2) //output [3 2 1]
```
