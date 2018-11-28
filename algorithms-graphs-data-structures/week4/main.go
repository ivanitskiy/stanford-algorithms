package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var (
		filePath string
		scanner  *bufio.Scanner
		array    []int
	)

	flag.StringVar(&filePath, "f", "text.txt", "Path to a file. If not provided then from stdin")
	flag.Parse()
	data := make(map[int][]int)
	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()
	scanner = bufio.NewScanner(bufio.NewReader(f))
	i := 0
	array = make([]int, 1000000)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		checkError(err)
		array[i] = number
		data[number] = []int{-10000 - number, 10000 - number}
		i++
	}
	hash := make(map[int]int)
	retval := make(map[int]int)
	i = 0
	for x := range data {
		if _, ok := hash[x]; !ok {
			if _, ok := hash[-10000-x]; ok {
				break
			}
			if _, ok := hash[10000-x]; ok {
				break
			}

			for y := -10000 - x; y <= 10000-x; y++ {
				if _, ok := data[y]; ok {
					hash[y] = x
					hash[x] = y
					retval[x+y] = x + y
					// fmt.Println("found x,y", x, y, x+y)
				}
			}
			// fmt.Println(len(retval), i, "/", len(data))
		}
		i++
	}

	// sort.Ints(array)

	fmt.Println(len(data), len(retval))

}
