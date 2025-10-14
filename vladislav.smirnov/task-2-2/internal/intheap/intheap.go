package intheap

type IntHeap []int

func (h *IntHeap) Len() int {
  return len(*h)
}

func (h *IntHeap) Swap(i, j int) {
  (*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Less(i, j int) bool {
  return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Push(x any) {
  value, ok := x.(int)
  if !ok {
    panic("bad cast to int")
  }
  *h = append(*h, value)
}

func (h *IntHeap) Pop() any {
  if h.Len() == 0 {
    return nil
  }

  x := (*h)[len(*h)-1]
  *h = (*h)[:len(*h)-1]
  return x
}
