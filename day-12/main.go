package main

import (
	"os"
	"strings"
)

func parseGraph(lines []string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range lines {
		chs := strings.Split(line, "-")
		from, to := chs[0], chs[1]
		if _, ok := graph[from]; !ok {
			graph[from] = make([]string, 0, 1)
		}
		graph[from] = append(graph[from], to)
		if _, ok := graph[to]; !ok {
			graph[to] = make([]string, 0, 1)
		}
		graph[to] = append(graph[to], from)
	}
	return graph
}

func computeOffsetMap(graph map[string][]string) map[string]int {
	offset := make(map[string]int)
	ix := 0
	for k, vs := range graph {
		if _, ok := offset[k]; !ok {
			offset[k] = ix
			ix++
		}
		for _, v := range vs {
			if _, ok := offset[v]; !ok {
				offset[v] = ix
				ix++
			}
		}
	}
	return offset
}

func traverseAll(graph map[string][]string) [][]string {
	offset := computeOffsetMap(graph)
	var visit func(string, uint64) [][]string
	visit = func(node string, visited uint64) [][]string {
		if visited&(1<<offset[node]) > 0 {
			return nil
		}
		if node == "end" {
			return [][]string{{node}}
		}
		res := make([][]string, 0, 1)
		if node[0] >= 'a' && node[0] <= 'z' {
			visited |= (1 << offset[node])
		}
		for _, nb := range graph[node] {
			for _, vis := range visit(nb, visited) {
				res = append(res, append([]string{node}, vis...))
			}
		}
		return res
	}
	return visit("start", 0)
}

func traverseAllWithRep(graph map[string][]string) [][]string {
	offset := computeOffsetMap(graph)
	/*
		visited = xx xx xx 00 xx xx xx xx
		visited = xx xx xx 01 xx xx xx xx
		visited = xx xx xx 11 xx xx xx xx
	*/
	var visit func(string, uint64, bool) [][]string
	visit = func(node string, visited uint64, canRep bool) [][]string {
		vv := (visited >> (2 * offset[node])) & 0b11
		if vv > 1 {
			return nil
		} else if vv > 0 {
			if node == "start" || node == "end" || !canRep {
				return nil
			}
		}
		if node == "end" {
			return [][]string{{node}}
		}
		res := make([][]string, 0, 1)
		rep := canRep
		if node[0] >= 'a' && node[0] <= 'z' {
			vv = (vv << 1) | 1
			visited |= (vv << (2 * offset[node]))
			rep = rep && (vv <= 1)
		}
		for _, nb := range graph[node] {
			for _, vis := range visit(nb, visited, rep) {
				res = append(res, append([]string{node}, vis...))
			}
		}
		return res
	}
	return visit("start", 0, true)
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	graph := parseGraph(lines)

	printf("graph: %+v", graph)

	traverses := traverseAll(graph)
	printf("traverses: %+v", traverses)

	printf("the result is: %d", len(traverses))

	traverses2 := traverseAllWithRep(graph)
	printf("traverses2: %+v", traverses2)

	printf("the result2 is: %d", len(traverses2))
}
