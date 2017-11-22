package main

import(
    "fmt"
    "sync"
    "time"
    "runtime"
)

var wg sync.WaitGroup

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
    wg.Add(2)
    go foo()
    go bar()
    wg.wait()
}

func foo() {

    defer wg.done()

    for i = 0; i < 100; i++ {
        fmt.Println("Foo = ", i)
        time.Sleep(time.Duration(4 * time.Millisecond))
    }
}

func bar() {

    defer wg.done()

    for i = 0; i < 100; i++ {
        fmt.Println("Boo = ", i)
        time.Sleep(time.Duration(20 * time.Millisecond))
    }
    wg.done()
}
