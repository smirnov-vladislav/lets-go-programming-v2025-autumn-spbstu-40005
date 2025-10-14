package main

import (
  "fmt"
  "container/heap"

  "task2_2/internal/intheap"
)

func main() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered. Error:\n", r)
    }
  }()

  var dishesCount int

  if _, err := fmt.Scan(&dishesCount); err != nil {
    fmt.Println("error scanning count of dishes: &v\n", err)
    return
  }

  dishes := &intheap.IntHeap{}
  heap.Init(dishes)

  for range dishCount {
    val value int
    if _, err := fmt.Scan(&value); err != nil {
      fmt.Println("error scanning value of dish: %v\n", err)
      return
    }

    heap.Push(dishes, value)
  }

  var wished int
  if _, err := fmt.Scan(&wished); err != nil {
    fmt.Println("error scanning wished dish: %v\n", err)
    return
  }

  for range wished - 1 {
    if heap.Pop(dishes) == nil {
      fmt.Println("error: heap is empty")
      return
    }
  }

  bestDish := heap.Pop(dishes)
  if bestDish == nil {
    fmt.Println("error: heap is empty")
    return
  }

  fmt.Println(bestDish)
}

