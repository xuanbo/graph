package graph

import "testing"

// TestNewDag 测试创建图
func TestGraph(t *testing.T) {
	draph := NewGraph()
	t.Log(draph.AddVertex("a"))
	t.Log(draph.AddVertex("b"))
	t.Log(draph.AddVertex("c"))
	t.Log(draph.AddVertex("d"))
	t.Log(draph.AddVertex("e"))
	t.Log(draph.AddVertex("f"))
	t.Log(draph.AddVertex("g"))
	t.Log(draph.AddVertex("h"))
	t.Log(draph.AddVertex("i"))
	t.Log(draph.AddVertex("j"))
	// 默认初始化10个顶点，此处扩容
	t.Log(draph.AddVertex("k"))
	t.Log(draph.AddEdge("a", "b"))
	t.Log(draph.AddEdge("b", "c"))
	t.Log(draph.AddEdge("c", "d"))
	t.Log(draph.AddEdge("d", "a"))
	// isDAG
	t.Logf("isDAG: %v", draph.IsDAG())
	t.Log(draph.AddEdge("i", "j"))
	// isDAG
	t.Logf("isDAG: %v", draph.IsDAG())
}
