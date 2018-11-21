package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds element into the collection
func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop removes item from the collection
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An MaxHeap is a max-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds element into the collection
func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop removes item from the collection
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type sequence struct {
	max *MaxHeap
	min *MinHeap
}

// newSequence creates a datastructure that supports insert and median operations
func newSequence() *sequence {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	maxHeap := &MaxHeap{}
	heap.Init(minHeap)

	s := &sequence{max: maxHeap, min: minHeap}
	return s
}

func (s *sequence) Insert(x int) {
	// insert the first element into maxHeap
	if len(*s.max) == 0 {
		heap.Push(s.max, x)
		return
	}

	if x > (*s.max)[0] {
		heap.Push(s.min, x)
	} else {
		heap.Push(s.max, x)
	}

	// Step 2: Balance the heaps (after this step heaps will be either balanced or
	// one of them will contain 1 more item)

	diff := len(*s.min) - len(*s.max)
	if diff > 0 {
		v := heap.Pop(s.min)
		heap.Push(s.max, v)
	} else if diff < 0 {
		v := heap.Pop(s.max)
		heap.Push(s.min, v)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (s *sequence) Median() int {
	var retval int
	var maxheaproot, minheaproot int
	if len(*s.min) > 0 {
		minheaproot = (*s.min)[0]
	}
	if len(*s.max) > 0 {
		maxheaproot = (*s.max)[0]
	}
	if len(*s.min) > len(*s.max) {
		retval = minheaproot
	} else {
		retval = maxheaproot
	}
	return retval
}

type sequencer interface {
	Insert(int)
	Median() int
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	defer timeTrack(time.Now(), "Computing Median Maintance")
	var filePath string
	seq := newSequence()

	flag.StringVar(&filePath, "f", "text.txt", "Path to a file. If not provided then from stdin")
	flag.Parse()
	var scanner *bufio.Scanner
	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()
	scanner = bufio.NewScanner(bufio.NewReader(f))
	retval := 0
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		checkError(err)
		seq.Insert(number)
		retval += seq.Median()
	}
	fmt.Println(retval)
}
