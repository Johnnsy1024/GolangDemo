package sortAlgorithm

func BubbleSort(ll []int) {
	n := len(ll)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if ll[j] > ll[j+1] {
				ll[j], ll[j+1] = ll[j+1], ll[j]
			}
		}
	}

}
