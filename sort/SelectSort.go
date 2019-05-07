package sort


func SelectSort(s []int) {
	if len(s) <= 1{
		return
	}
	for i:= range s {
		min := s[i]
		for j :=i; j < len(s); j++ {
			if s[j] < min {
				s[j], min = min, s[j]
			}
		}
		s[i] = min
	}
}