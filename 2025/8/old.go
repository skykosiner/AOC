package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/skykosiner/AOC/pkg/utils"
)

type JunctionBox struct {
	x, y, z int
}

type Edge struct {
	p1, p2 int
	dist   int
}

type UnionFind struct {
	parent, size []int
	count        int
}

func NewUnionFind(n int) UnionFind {
	uf := UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
		count:  n,
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Find(p int) int {
	if uf.parent[p] == p {
		return p
	}

	uf.parent[p] = uf.Find(uf.parent[p])
	return uf.parent[p]
}

func (uf *UnionFind) Union(p, q int) bool {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	if rootP != rootQ {
		if uf.size[rootP] < uf.size[rootQ] {
			rootP, rootQ = rootQ, rootP
		}

		uf.parent[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
		uf.count--

		return true
	}

	return false
}

func (uf *UnionFind) GetCircuitSizes() []int {
	sizes := make(map[int]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		sizes[root] = uf.size[root]
	}

	result := []int{}
	for _, size := range sizes {
		result = append(result, size)
	}

	return result
}

func NewJunctionBox(line string) JunctionBox {
	var box JunctionBox
	items := strings.Split(line, ",")

	x, _ := strconv.Atoi(items[0])
	y, _ := strconv.Atoi(items[1])
	z, _ := strconv.Atoi(items[2])

	box.x = x
	box.y = y
	box.z = z

	return box
}

func dist(point1, point2 JunctionBox) int {
	dx := point2.x - point1.x
	dy := point2.y - point1.y
	dz := point2.z - point1.z

	return dx*dx + dy*dy + dz*dz
}

func aoeu() {
	test := flag.Bool("test", false, "Run test case or not")
	flag.Parse()

	var inputFile string
	if *test {
		inputFile = "./input.test"
	} else {
		inputFile = "./input"
	}

	lines := utils.ReadFile(inputFile)

	boxes := []JunctionBox{}
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			boxes = append(boxes, NewJunctionBox(line))
		}
	}

	var edges []Edge
	for i := range boxes {
		for j := 1; j < len(boxes); j++ {
			distance := dist(boxes[i], boxes[j])
			edges = append(edges, Edge{i, j, distance})
		}
	}

	fmt.Println(edges)

	sort.SliceStable(edges, func(i, j int) bool {
		if edges[i].dist != edges[j].dist {
			return edges[i].dist < edges[j].dist
		}

		// sort by p1 index
		if edges[i].p1 != edges[j].p1 {
			return edges[i].p1 < edges[j].p1
		}

		// sort by p2 index
		return edges[i].p2 < edges[j].p2
	})

	uf := NewUnionFind(len(boxes))
	connections := 0
	edgeIdx := 0

	targetConnections := 1000
	if *test {
		targetConnections = 10
	}

	for connections < targetConnections && edgeIdx < len(edges) {
		edge := edges[edgeIdx]
		if uf.Union(edge.p1, edge.p2) {
			connections++
		}
		edgeIdx++
	}

	sizes := uf.GetCircuitSizes()
	sort.Ints(sizes)
	fmt.Println(sizes)
	largestSizes := sizes[len(sizes)-3:]
	sum := largestSizes[0] * largestSizes[1] * largestSizes[2]
	fmt.Println(sum)
}
