package graph

import (
	"container/list"
	"fmt"
)

type graph struct {
	cap int
	adj []*list.List
}

func NewGraph(cap int) *graph {
	g := &graph{cap:cap}
	g.adj = make([]*list.List, cap)
	for i := range g.adj {
		g.adj[i] = list.New()
	}
	return g
}

func (g *graph) AddEdge(s, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}


func (g *graph) BFS(s, t int) {
	if s == t {
		return
	}
	// 前置节点记录列表,先全部初始化为-1
	prev := make([]int, g.cap)
	for i := range prev{
		prev[i] = -1
	}
	// 记录待遍历的节点
	queue := []int{s}
	// 记录访问过的节点
	visited := make([]bool, g.cap)
	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linklist := g.adj[top]
		for e := linklist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}
	if isFound {
		prinPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}


}

func (g *graph)recurse(s, t int, prev []int, visited []bool, isFound bool) {
	if isFound {
		return
	}
	visited[s] = true
	if s == t {
		isFound = true
		return
	}
	linklist := g.adj[s]
	for e := linklist.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurse(k, t, prev, visited, isFound)
		}
	}
}

func (g *graph) DFS(s, t int) {
	var prev = make([]int, g.cap)
	for i := range prev {
		prev[i] = -1
	}
	visited := make([]bool, g.cap)
	isFound := false
	g.recurse(s, t, prev, visited, isFound)
	if isFound {
		prinPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}


func prinPrev(prev []int, s, t int) {
	if s == t || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		prinPrev(prev, s, prev[t])
		fmt.Printf("%d\n", t)
	}
}
