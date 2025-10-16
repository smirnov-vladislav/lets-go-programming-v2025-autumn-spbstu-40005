package main

import (
  "fmt"
  "container/heap"

  "github.com/smirnov-vladislav/task-2-2/internal/intheap"
)

func main() {

  var dishesCount int

  if _, err := fmt.Scan(&dishesCount); err != nil {
    fmt.Printf("error scanning count of dishes: %v\n", err)

    return
  }

  dishes := &intheap.IntHeap{}
  heap.Init(dishes)

  for range dishesCount {
    var value int
    if _, err := fmt.Scan(&value); err != nil {
      fmt.Printf("error scanning value of dish: %v\n", err)

      return
    }

    if _, err := heap.Push(dishes, value); err != nil {
      fmt.Println("heap push error:", err)
    }
  }

  var wished int
  if _, err := fmt.Scan(&wished); err != nil {
    fmt.Printf("error scanning wished dish: %v\n", err)

    return
  }

  for range wished - 1 {
    if heap.Pop(dishes) == nil {
      fmt.Printf("error: heap is empty")

      return
    }
  }

  bestDish := heap.Pop(dishes)
  if bestDish == nil {
    fmt.Printf("error: heap is empty")

    return
  }

  fmt.Println(bestDish)
}
