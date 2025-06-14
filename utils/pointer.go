package utils

func ChangeSlice(s []int) {
	s = []int{100, 200, 300}
}

func ChangeSlice2(s *[]int) {
	*s = []int{100, 200, 300}
}
