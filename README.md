# Golang

- [Golang](#golang)
  - [Installation and Setup](#installation-and-setup)
    - [Install golang](#install-golang)
    - [IDE setup](#ide-setup)
  - [Create main func](#create-main-func)
  - [Declare variables and functions](#declare-variables-and-functions)
  - [array and slices](#array-and-slices)
  - [Go receiver and function argument](#go-receiver-and-function-argument)
  - [go data type](#go-data-type)
    - [type convertion](#type-convertion)
    - [struct](#struct)
    - [Pointers](#pointers)

## Installation and Setup

### Install golang

Download go installation package from [golang], follow the instruction of **Download Go**.  

In your terminal, use `go help` to make sure you have installed go env successfully.

### IDE setup

[vscode][vscode link] can be a good ide for golang programming.
once installed vscode, go to `Go`->`Extension`, install one of golang editing plugin, which can support compile checking.

## Create main func

`main.go` is the entrypoint for go program. Create a main.go lang file in your code directory.

Here is a very simple hello world code of go lang.

    package main

    import "fmt"

    func main() {
        fmt.Println("Hello World)
    }
You can run thie code both in vscode by `Run`->`Start Debugging` or from terminate, by `go run main.go`

## Declare variables and functions

There are ways of creating go variables, below are some ways of creating go variables

    var myvar string = "my string" // define myvar as string type, and assign value "my string" to it
    myvar := "my string" // go lang will automatically figure out the type of myvar is string, and assign the value to it
    c := make(chan string) // use build-in make function to create a channel of type string.

Funcation can be defined as below

    func myfunc() {
        // a simple function without returning data
    }

    func myfunc2() string {
        // a func return string
        return "my string"
    }

    func myfunc3() (string, int) {
        // a func with multiple return data
        return "my string", 1
    }

    func myfunc4(s string) int {
        // a func with pass in a string argument
        return len(s)
    }

    func (a mytype) myfunc5() {
        // a func receives data a belongs mytype, a can be used in func
        // Please be aware go only copy another data from a to the func, modifying data in func won't change the originally variable 
    }

    func (a *mytype) myfunc6() {
        // a func receives address of a variable, the func can modify a's properties
    }

## array and slices

In Go, array has fixed size. Slices add a dynamic layer on top of arrays.
Slices can be created without declaring the length, and build-in append() func can be used to add element in slices.  

    myColorSlices := []string{"red", "yellow"}
    myColorSlices = append(myColorSlices, "green")

Also developer can use index to return a subset set of the slides, e.g. `myColorSlices[start:end]`. Like some other language, start and end field can be left empty so that by default it will point to the beginning and end of slices.

Iterate slices  

    for idx, color := range myColorSlices {
        // ...
    }
you can replace idx or color to `_` if you don't expect to use it

## Go receiver and function argument

Go is not Object oriented language, there is no concept of object and variable and functions defined in object class as Java. Instead, go use `Receiver`, a function can receive copy of data or pointer of originally data as input. e.g.

To use this feature, firstly we need to define a go type. e.g.

    type myType []string
defined a type named `myType`, which is an slides of string.

    func (m myType) myFunc() {
        ...
    }

myFunc receives myType data(which is a copy of myType in different memory address here). `myFunc` can be called by following way

    myType.myFunc()

Another way is passing variable as argument to func  

    func deal(d deck, handSize int) (deck, deck) {
        return d[:handSize], d[handSize:]
    }

## go data type

### type convertion

go lang can use type convertion to convert a type to another type.
e.g.  

    []byte(d.toString())

### struct

struct is a value type in go, it can be used to define structured variables with multiple fields, e.g  

    type contactInfo struct {
        email   string
        zipCode int
    }

    type person struct {
        firstName string
        lastName  string
        contact   contactInfo
    }

a new variable can be created by

    newperson := person{
        "wang",
        "ba",
        "110",
    }

### Pointers

In go lang, there are 2 types of data, which are `value type` and `reference type`.  
**value type**: int,float,string,bool,structs
**reference type**: slices, maps, channels, pointers, functions

value type in receivers will be copied to another address in memory by golang. Thus, any modification will be based on the new copy of data without changing the original data. To pass the origin variable to function in Go, we need use pointer.

To use pointer, we need to defined the receiver receives pointer(address) of variable in function defintion. e.g.
```func (p *person) updateName(newFirstName string)```
then we can either use `&` to retrieve pointers of the value  

    personPointer := &person
    personPointer.updateName("Tom")

or let go lang automatically pass pointer of variable when calling function. It needs us define receiver receives points.

    person.updateName("Tom")

### map

golang use map to store key-value pairs. In map, type of keys and values are required to be as a unique static type accordingly. e.g.
`var myMap map[KeyType]valueType`

Data can be added to map by
`mymap[key] = value`
And be deleted from map by
`delete(myMap, key)`

### Interface

In Go lang, if a type implements all functions defined in interface type, then the type will be automatically promoted to that interface type. And it can be received by any functions receive the interface type. e.g

    type chineseBot struct{}
    type englishBot struct{}
    type bot interface {
        getGreeting() string
    }
    func (chineseBot) getGreeting() string {
        return "你好"
    }
    func (englishBot) getGreeting() string {
        return "Hi"
    }
    func printGreeting(b bot) {
        fmt.Println(b.getGreeting())
    }

in this case, since the chineseBot implements getGreeting() func, thus, chineseBot variable will be recognized as bot type and thus be able to call printGreeting() . Interface in golang is implicit, we don't have to define a type belong to a interface, but we need to have a type implement all functions defined in interface type.

Be aware that interface is not a concrete type but an interface type, which means we can not create a variable of interface type directly.

Since interface doesn't have actual value, thus it can only be passed to function as parameter in function definition, e.g.
`func printArea(s shape)`

### routine

Go routines is lightweight thread managed by the Go runtime. Comparing to thread in Java, it's more scallable and lightweight due to:

1. The use of growable segment stack, which reduce the waste of fixed stack space in machine.

2. Handle over sharing data to channel, avoid complex locking operation in programming code. And improve performance accordingly.

Compare to some event driven framework like Nodejs with high performance, go lang is less complex in code.

Routine can be created by `go` keyword following by function.  

    for _, link := range links {
        go checkLink(link)
    }

### Channel

In Go lang, channel is used to support communication between routine. To use it, firstly you will need to define the type of data in channel. A channel can only host data with predefined data type
`c := make(chan string)`

Channel can be used for bi-way communiciation.
To write data to channel
`c <- data`
To read from a channel
`data <- c` or calling `<- c` directly which will pass variable from channel to caller function.

`<- c` is a blocking call, it will be unblocked only after current routine reading a variable from the channel.

## Iterate func

In Go, just as Lambda in Java, go support creating function iterately. e.g

    for l := range c {
        go func(link string) {
            time.Sleep(5 * time.Second)
            checkLink(link, c)
        }(l)
    }
be careful about argument in the function and avoid using same reference in main routine in function call of sub routine.

## Packages

golang has many packages to support different functions, it's highly recommended to read [package page][go package] when you want to build up functions by other peoples wise.

[golang]: https://golang.org/
[vscode link]: https://code.visualstudio.com
[go package]: https://golang.org/pkg/

## project

In project, I wrote a web crawler by Go lang.