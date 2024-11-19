package main

import (
	"math/rand"
)

type RandomizedSet struct {
	// insert, deletion -> O(1)
	// random -> array, hashmap [val] -> ind O(1)
	mp  map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	rs := RandomizedSet{}
	rs.mp = make(map[int]int)
	rs.arr = make([]int, 0)
	return rs
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mp[val]; ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.mp[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.mp[val]; !ok {
		return false
	}
	ind := this.mp[val]
	sz := len(this.arr)
	this.mp[this.arr[sz-1]] = ind
	this.arr[ind], this.arr[sz-1] = this.arr[sz-1], this.arr[ind]
	this.arr = this.arr[:sz-1]
	delete(this.mp, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randInd := rand.Intn(len(this.arr))
	return this.arr[randInd]
}
