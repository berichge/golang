# Go lang

This is a memo of Go lang study

## Installation and Setup

Installation: go to https://golang.org/, follow the instruction of **Download Go**.  
Setup: [Visual studio code][vscode link] can be a good ide of golang programming.  

## Create main func

create a main.go file
package main

    import "fmt"

    func main() {
        fmt.Println("Hello World)
    }

## Declare variables and functions

There are 2 ways of initializating go lang variables, which are listed below. Go lang is statically-typed language, which performs type checking at compile time. We can either declare the type of variable or use **:=** to have go lang figure out type automatically during compile.  

    var myvar string = "my string"
    var myvar := "my string", for the first time initialization

Funcation can be defined as below, the last decorator indicates the return type of function.  

    func myfunc string() {
        return "my string"
    }

## Array and slices

In Go, array has a fixed size, which is a part of definition of an array. Slices add a dynamic layer on top of arrays. Slices can share data in same array, also a built-in function copy() can create a copy of slices.  
Slices can be created without declaring the length of array, and append() can be called to added element in slices.  

    myColor := []string{"red", "yellow"}
    myColor = append(myColor, "green")
And slides can use index to return a subset set of the slides, like *start:end*.

Iterate slices  

    for idx, color := range colors {
        // ...
    }
you can replace idx to **_** if you don't expect to use it

## Go receiver and argument

Go is not Object oriented language, thus it don't use object to call functions. Instead, go use **Receiver** to define how function process data. An example can be shown as below  

    type myTypeRef []string
    func myReceiver(m myType) {
        ...
    }
*m myType* is a receiver here, the first word is the reference of data, and the second one is the expected type
In Go lang, type was an wrapper of basic datatype, and func can be defined in same file with type if there logic are related. Receiver can only receive valid types which had been defined. `map[string]string` is not a valid type in receiver.

Another way is pass parameter as argument to func  

    func deal(d deck, handSize int) (deck, deck) {
        return d[:handSize], d[handSize:]
    }

## Return value

in the return annotation field, we can use (type1, type2) to let func return multiple value. e.g.
```hand, remainingDeck := deal(cards, 5)```

## Type convertion

go lang can use type convertion to convert type to another type.
e.g.  

    []byte(d.toString())

## struct

struct is a value type in go, it can be used to define structured variables with multiple fields
e.g  

    type contactInfo struct {
        email   string
        zipCode int
    }

    type person struct {
        firstName string
        lastName  string
        contact   contactInfo
    }

## Pointers

In go lang, there are 2 types of data, which are value type and reference type.  
**value type**: int,float,string,bool,structs
**reference type**: slices, maps, channels, pointers, functions

value type in receivers will be copied to another address in RAM by golang. Thus, any modification will be based on the new copy of data without changing the original data. To pass the reference of value to function in Go, we need use pointer.

There are ways of using pointer. But in the func receiver defintion, we need to defined the receiver receives point(address) of data at first. e.g.
```func (p *person) updateName(newFirstName string)```
then we can either use **&** to retrieve pointers of the value.  

    personPointer := &person
    personPointer.updateName("Tom")
Or go lang will automatically convert value be passed in function to pointer with above definition. 

    person.updateName("Tom")

## Map

Go lang use map to store key-value pairs. In map, type of keys and values are predefined as a unique static type accordingly. e.g.
`var mymap map[int]string`
go has build-in func delete to remove element from map.
`delete(map, "key")`

## Interface

In Go lang, if a type implements all functions defined in interface type, then the type will be automatically promoted to that interface type. And it can be received by any functions receive the interface type.  

    type chineseBot struct{}
    type bot interface {
        getGreeting() string
    }
    func (chineseBot) getGreeting() string {
        return "你好"
    }
    func printGreeting(b bot) {
        fmt.Println(b.getGreeting())
    }
in the case, since the chinese bot implements getGreeting() func, thus, chinesebot variable will be able to call printGreeting() func. Interface in go lang is implicit, we don't have to define a type belong to a interface.

Be aware that interface is not a concrete type but an interface type, which means we can not create a variable of interface type directly.

Since interface doesn't have actual value, thus it can only be pass to function as parameter, e.g.
`func printArea(s shape)`

## Packages

Use packages for basic operation
[strings][go strings package]
[ioutil][go io package]
[rand][rand package]

[vscode link]: https://code.visualstudio.com
[go strings package]: https://golang.org/pkg/strings/
[go io package]: https://golang.org/pkg/io/ioutil/
[rand package]: https://golang.org/pkg/math/rand
