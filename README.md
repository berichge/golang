# Go lang

This is a memo of Go lang study

## Installation and Setup

Installation: go to https://golang.org/, follow the instruction of **Download Go**.  
Setup: [Visual studio code][vscode link] can be a good ide of golang programming.  

[vscode link]: https://code.visualstudio.com

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
    var myvar := "my string"

Funcation can be defined as below, the last decorator indicates the return type of function.  

    func myfunc string() {
        return "my string"
    }

## Array and slices

In Go, array has a fixed size, which is a part of definition of an array. Slices add a dynamic layer on top of arrays. Slices can share data in same array, also a built-in function copy() can create a copy of slices.  
Slices can be created without declaring the length of array, and append() can be called to added element in slices.  

    myColor := []string{"red", "yellow"}
    myColor = append(myColor, "green")
