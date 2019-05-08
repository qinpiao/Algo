package WeightRandom


import (
	"testing"
	"fmt"
)

func TestWeightRandom(t *testing.T) {
	s := map[interface{}]int{"A": 10, "B": 2, "C": 3, "D":5}
	for i:= 0; i< 20; i++ {
		fmt.Println(WeightRandom(s))
	}

}
