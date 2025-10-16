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
    return fmt.Errorf("bad cast to int: %v", x)
  }
  *h = append(*h, value)

  return nil
}

func (h *IntHeap) Pop() any {
  if h.Len() == 0 {
    return nil
  }

  orig := *h
  new := len(old)
  x := orig[n-1]
  *h = orig[0 : n-1]

  return x
}
