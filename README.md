# Go

- Create a go module
- the name of the module should be a valid. It can also be a valid url path.
- as soon as the module is created , it creates a file called go.mod
- go.mod file contains dependencies as well. 

```go mod init <name of the module>```

- Notes
- there are two types of functions.
- 1. builtin 
- 2. functions from packages.

- there are three kinds of pacakges
- 1. stadard packages
- 2. user defined packages
- 3. third party packages

## builtin
- append,cap,copy,delete,len,make,new,panic,print,println

## keywords

- break,case,const,continue,default,defer,else,fallthrough,for,func,goto,if,import,map,package,range,return,struct,switch,type,var

## compile build run

- run 

```
go run hello.go 
go run . # it automatically looks for main func
go run --work hello.go
```

## built and generate binary or exe

```
go build main.go
```

- build with output file name

```
go build -o demo main.go
```
- to stripedown the binary

```
go build -ldflags "-w" -o demo main.go
```

- to list out all distributions

```
go tool dist list
```

- to cross compile and build

    1. change GOOS and GOARCH environment variables for respective os/arch

```
GOOS="darwin" GOARCH="amd64" go build -o apple-demo main.go
```