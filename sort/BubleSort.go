package sort

func BubleSort(s []int){
	if len(s) <= 1 {
		return
	}
	stop := len(s)-1
	for stop > 0 {
		for i := range s[:stop] {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
		stop--
	}
}

func BubleSort2(s []int){
	if len(s) <= 1 {
		return
	}
	stop := len(s)-1
	swap := true
	for stop > 0 {
		if !swap {
			break
		}
		for i := range s[:stop] {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swap = true
			} else {
				swap = false
			}
		}
		stop--
	}
}
