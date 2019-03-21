package main

import (
	"fmt"
	"math"
)

type IntSet struct {
	set map[int64]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[int64]bool)}
}

func (set *IntSet) Add(i int64) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found // false if already exist
}

func (set *IntSet) length() int {
	return len(set.set)
}

func power_number() int {
	power_set := NewIntSet()
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			a_pow_b := math.Pow(float64(a), float64(b))
			power_set.Add(int64(a_pow_b))
		}
	}
	return power_set.length()
}

func main() {
	res := power_number()
	fmt.Printf("res=%d\n", res)
}
