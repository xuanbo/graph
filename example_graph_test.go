package graph_test

import (
	"fmt"

	"github.com/xuanbo/graph"
)

func Example() {
	draph := graph.NewGraph()
	// 添加顶点
	draph.AddVertex("a")
	draph.AddVertex("b")
	draph.AddVertex("c")
	draph.AddVertex("d")
	draph.AddVertex("e")
	draph.AddVertex("f")
	draph.AddVertex("g")
	draph.AddVertex("h")
	draph.AddVertex("i")
	draph.AddVertex("j")
	// 默认初始化10个顶点，此处扩容一次
	draph.AddVertex("k")
	// 添加边
	draph.AddEdge("a", "b")
	draph.AddEdge("b", "c")
	draph.AddEdge("c", "d")
	draph.AddEdge("d", "a")
	// isDAG
	fmt.Printf("isDAG: %v\n", draph.IsDAG())
}
