package main

import (
	"container/list"
	"fmt"
)

type Graph struct {//无向图
	v int
	adj []*list.List
}

func New(v int) *Graph  {
	graph := &Graph{v,make([]*list.List,v)}
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

func (g *Graph) addEdge(s , t int)  {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

//广度优先
func (g *Graph) BFS(s , t int)  {
	if s == t {
		return
	}
	//已访问过
	visited := make([]bool , g.v)
	visited[s] = true
	//存储一访问节点，但相连节点未访问过
	queue := make([]int , g.v)
	queue[0] = s
	//存储一访问节点的路径
	prev := make([]int , g.v)
	for i:= range prev {
		prev[i] = -1
	}
	for len(queue) != 0 {
		w := queue[0]
		queue = queue[1:]
		list := g.adj[w]
		for e := list.Front(); e != nil; e = e.Next() {
			q := e.Value.(int)
			if !visited[q] {
				prev[q] = w
				if q == t {
					g.Print(prev , s ,t)
					return
				}
				visited[q] = true
				queue = append(queue,q)
			}
		}

	}
}


var found bool

func (g *Graph) Print(prev []int , s ,t int)  {
	if prev[t] != -1 && t != s {
		g.Print(prev , s , prev[t])
	}
	fmt.Printf("%d->",t)
}



//深度优先
func (g *Graph) DFS(s , t int)  {
	found = false
	//已访问过
	visited := make([]bool , g.v)
	//存储一访问节点的路径
	prev := make([]int , g.v)
	for i:= range prev {
		prev[i] = -1
	}
	g.recurDfs(s,t,visited,prev)
	g.Print(prev,s,t)

}

func (g *Graph) recurDfs(w , t int , visited []bool , prev []int){
	if found {
		return
	}
	visited[w] = true
	if w == t {
		found = true
		return
	}
	list := g.adj[w]
	for e := list.Front(); e != nil; e =e.Next()  {
		q := e.Value.(int)
		if !visited[q] {
			prev[q] = w
			g.recurDfs(q , t , visited , prev)
		}
	}

}



func main()  {
	graph := New(10)
	graph.addEdge(1,2)
	graph.addEdge(1,4)
	graph.addEdge(2,4)
	graph.addEdge(2,3)
	graph.addEdge(2,5)
	graph.addEdge(5,3)

	graph.BFS(1,5)
	fmt.Println()
	graph.DFS(1,5)

}


