type Item struct {
	Val int
	Frequency int
}
type MinHeap []Item
func (h MinHeap) Len() int { return len(h)}
func (h MinHeap) Less(i, j int) bool { return h[i].Frequency < h[j].Frequency }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, freq := range freqMap {
		heap.Push(h, Item{Val: num, Frequency: freq})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)

	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(Item).Val)
	}

	return res

}
