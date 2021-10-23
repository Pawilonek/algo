# Pawilonek's alog

⚠️ My implementation of some algorithms. This is a pure learning repository so if you want to use it do it at your own risk.

# Algorithms

## Queue

FIFO queue that accepts any type of variable.
Example usage:
```go

import "github.com/pawilonek/algo/queue"

[..]

q := queue.NewQueue()
q.Push("element 1")
q.Push("element 2")

for !q.IsEmpty {
    fmt.Println(q.Pop())
}
```

## 
