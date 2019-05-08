package WeightRandom

import (
	"math/rand"
	"time"
)


func WeightRandom(s map[interface{}]int) interface{} {
	rand.Seed(time.Now().UnixNano())
	edge := 0
	for _, v := range s {
		edge += v
	}
	c := rand.Intn(edge)
	l := 0
	for k, v := range s {
		l += v
		if l > c {
			return k

		}
	}
	return nil
}