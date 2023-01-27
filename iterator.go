package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type element interface{}

type weightFunc func(element) int

type iterator interface {
	next() bool
	val() element
}

type intIterator struct {
	idx   int
	value []element
}

func (intIt *intIterator) next() bool {
	if intIt.idx < len(intIt.value) {
		return true
	}
	return false
}

func (intIt *intIterator) val() element {
	if intIt.next() {
		val := intIt.value[intIt.idx]
		intIt.idx++
		return val
	}
	return nil
}

func newIntIterator(src []int) *intIterator {

	iterator := &intIterator{idx: 0, value: make([]element, len(src))}
	for i := range src {
		iterator.value[i] = src[i]
	}
	return iterator
}

func main() {
	nums := readInput()
	it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}

	m := max(it, weight)
	fmt.Println(m)

}

func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
