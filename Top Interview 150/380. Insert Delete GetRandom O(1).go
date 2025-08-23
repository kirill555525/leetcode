package main

import "math/rand"

type RandomizedSet struct {
	m map[int]int

	s []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		s: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.s = append(this.s, val)
		this.m[val] = len(this.s) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.m[val]; ok {
		indexLast := len(this.s) - 1
		last := this.s[indexLast]
		this.s[i] = last
		this.m[last] = i
		this.s = this.s[:indexLast]
		delete(this.m, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	n := len(this.s)
	return this.s[rand.Intn(n)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
