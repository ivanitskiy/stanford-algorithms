// The file contains the adjacency list representation of a simple undirected graph. There
// are 200 vertices labeled 1 to 200. The first column in the file represents the vertex label,
// and the particular row (other entries except the first column) tells all the vertices that the
// vertex is adjacent to. So for example, the 6th row looks like : "6	155	56	52	120	......".
// This just means that the vertex with label 6 is
// adjacent to (i.e., shares an edge with) the vertices with labels 155,56,52,120,......,etc

// Your task is to code up and run the randomized contraction algorithm for the min cut problem
// and use it on the above graph to compute the min cut. (HINT: Note that you'll have to figure out
// an implementation of edge contractions. Initially, you might want to do this naively,
// creating a new graph from the old every time there's an edge contraction. But you should also
// think about more efficient implementations.) (WARNING: As per the video lectures, please make
// sure to run the algorithm many times with different random seeds, and remember the smallest cut
// that you ever find.) Write your numeric answer in the space provided. So e.g.,
// if your answer is 5, just type 5 in the space provided.
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type graph struct {
	edges         map[string][]string
	lastNodeLabel int
}

func (g *graph) addEdge(from, to string) {
	g.touchNode(to)
	g.touchNode(from)
	g.edges[to] = append(g.edges[to], from)
	g.edges[from] = append(g.edges[from], to)
}

// returns a list of nodes with labels
func (g *graph) nodes() []string {
	nodes := make([]string, 0, len(g.edges))
	for k := range g.edges {
		nodes = append(nodes, k)
	}
	return nodes
}
func (g *graph) removeEdge(to, from string) {

	deleteValue := func(edges []string, value string) []string {
		index := -1
		for i, v := range edges {
			if v == value {
				index = i
			}
		}
		if index > -1 {
			edges = append(edges[:index], edges[index+1:]...)
		}
		return edges
	}
	g.edges[from] = deleteValue(g.edges[from], to)
	g.edges[to] = deleteValue(g.edges[to], from)
}

func (g *graph) mergeNodes(to, from string) {
	newLabel := fmt.Sprintf("%v", g.lastNodeLabel+1)
	g.touchNode(newLabel)
	newAdjacents := append(g.edges[to], g.edges[from]...)
	g.edges[newLabel] = newAdjacents
	copy(g.edges[newLabel], newAdjacents)
	delete(g.edges, to)
	delete(g.edges, from)

	for label, adjacents := range g.edges {
		for i, v := range adjacents {
			if v == to || v == from {
				adjacents[i] = newLabel
			}
		}
		b := adjacents[:0]
		for _, x := range adjacents {
			if x != label {
				b = append(b, x)
			}
		}
		g.edges[label] = b
	}
}

func (g *graph) touchNode(n string) {
	if _, ok := g.edges[n]; !ok {
		g.edges[n] = make([]string, 0)
		g.lastNodeLabel++
	}
}

func newGraph() *graph {
	g := &graph{edges: make(map[string][]string), lastNodeLabel: 0}
	return g
}

func fromFile(path string) (*graph, error) {
	data, err := ioutil.ReadFile(path)

	g := newGraph()

	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		nodes := strings.Split(line, "\t")
		g.touchNode(nodes[0])
		g.edges[nodes[0]] = nodes[1:]
	}
	return g, nil
}

func randomEdge(g graph) (string, string) {

	rand.Seed(time.Now().UnixNano())

	randNumber := rand.Intn(len(g.nodes()))
	src := g.nodes()[randNumber]
	randNumber = rand.Intn(len(g.edges[src]))
	dst := fmt.Sprintf("%v", g.edges[src][randNumber])
	return src, dst
}

func main() {

	retval := 10000000
	iter := 0
	for i := 0; i < 1000; i++ {
		g, _ := fromFile("kragerMinCut.txt")
		for len(g.edges) > 2 {
			to, from := randomEdge(*g)
			g.mergeNodes(to, from)
			// fmt.Println("Merge completed:", len(g.edges))
		}
		// fmt.Println(*g)
		cuts := len(g.edges[g.nodes()[0]])
		if cuts < retval {
			retval = cuts
			iter++
		}
		if cuts == 17 {
			break
		}
		// g = nil
	}
	fmt.Println(retval, iter)
}
