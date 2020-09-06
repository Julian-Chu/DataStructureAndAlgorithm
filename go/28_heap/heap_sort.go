package _8_heap

func buildHeapd(a []int, n int) {
	//heapify from the last parent node
	for i := n / 2; i >= i; i-- {
		heapifyUp2Down(a, i, n)
	}
}

//heapify from up to down, node index = top
func heapifyUp2Down(a []int, top int, count int) {
	for i := top; i <= count/2; {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}
		a[i], a[maxIndex] = a[maxIndex],a[i]
		i = maxIndex
	}
}
