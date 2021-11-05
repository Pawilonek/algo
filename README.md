# Pawilonek's algo

⚠️ My implementation of some algorithms. This is a pure learning repository so if you want to use it do it at your own risk.

# Instalation

Instal as a module for your project:

```bash
go get github.com/Pawilonek/algo
```

# Algorithms

## Queue

FIFO queue that accepts any type of variable. Example usage:

```go

import "github.com/Pawilonek/algo/queue"

[..]

q := queue.NewQueue()
q.Push("element 1")
q.Push("element 2")

for !q.IsEmpty() {
    fmt.Println(q.Pop())
}
```

## Stack

FILO queue that accepts any type of variable. Example usage:

```go

import "github.com/Pawilonek/algo/stack"

[..]

s := stack.New()
s.Push("element 1")
s.Push("element 2")

for !s.IsEmpty() {
    fmt.Println(s.Pop())
}
```

