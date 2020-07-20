## Go Basic

### Defer

A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. (***Easy variables cleanup actions***)

Allow us think about closing each file right after opening it.

Rules of ***Defer*** functions

- A deferred function's arguments are evaluated when the defer statement is evaluated.
- Deferred functions are executed in last in first out order.
- Deferred functions may read and assign to the returning function's named return values.

It can be called with function defintion
```
defer func() {
    xxx
}()
```

#### Use of defer

```
mu.Lock()
defer mu.Unlock()
```

### Panic

When function F calls ***panic***, execution of F stops, any refrred functions in F are executed normally, and then F returns to its caller.

It can be called by invoking panic directly, or can be caused by runtime errors.

### Recover

A built-in function that regains control of a panicking goroutine. Useful inside referred functions. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

### If 

Basic syntax

```
if x > max {
    x = max
} else {
    ...
}
```

with init statement
```
if x := f(); x <= y {
    return x
} else if x > z{
    ...
}
```

### []string vs ...string


