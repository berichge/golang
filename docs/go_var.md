# Go Variable

## Basic rules

- Golang is typed language, when variable are declared, they will be either explicitly or implicitly assigned a type.
- Golang requires every variable declare inside main() get used.
- Var has scope, start from defintion from a {, and end with }.

## Variable declaration and init

```
var i int   // declaration
var i int = 10 // declaration with init
var i = 10 // declaration omit types
name := "Donald J Poop" // short variable declaration
var name, age = "Donald J Poop", 10 // multiple declaration
// declare in bracket
if x {
    y :=1
}
// Variable declaration block, can improve code quality
var (
    var num1 int
    var num2 int
    var name string
)
```

## Data structure

### map

declaration: 
map[keyType]ValueType

`var m map[string]int`

#### init a map

`m = make(map[string]int)`

#### use map

```
m["a"] = 1
i := map["a"]
n := len(m)
delete(m, "route")
i, ok := m["a"]
_, ok := m["a"]
for k, v := range m {
    xxx
}
m = map[string]int{}
```

#### concurrency

    var counter = struct{
        sync.RWMutex
        m map[string]int
    }{m: make(map[string]int)}

    counter.RLock()
    n := counter.m["some_key"]
    counter.RUnlock()
    fmt.Println("some_key:", n)

    counter.Lock()
    counter.m["some_key"]++
    counter.Unlock()   

### Go interface

Specific a set of one or more methods signature.

```
resp, ok := callResp.RealResponse().(*AResponse)
```

Since callResp.RealResponse() return interface, the statement will check if callResp.RealResponse() return instance with type = AResponse, if yes, ***ok*** will be ***true***.

### Go Enum

lota represents successive integer constants, 0,1,2, it can also be 

```
const {
    c0 EnumTypeName = iota
    c1
    c2
}
```

Create a new integer type, list its values using iota.