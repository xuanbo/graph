package graph

const (
	// InitVertexNum 图的初始化顶点数
	InitVertexNum = 10
)

// Graph 图，邻接矩阵实现
type Graph struct {
	// 顶点表
	Vertex []string
	// 边
	Edge [][]int
	// 顶点数
	VerticeNum int
	// 边数
	EdgeNum int
	// 是否为有向无环图
	isDAG bool
	// 标记矩阵，0为当前顶点未访问，1为访问过，-1表示当前顶点访问过其他顶点。
	color []int
}

// IndexOfVertex 返回顶点名vertex的索引，不存在则返回-1
func (g *Graph) IndexOfVertex(vertex string) int {
	length := cap(g.Vertex)
	for i := 0; i < length; i++ {
		if vertex == g.Vertex[i] {
			return i
		}
	}
	return -1
}

// AddVertex 新增顶点vertex，如果顶点存在则返回false
func (g *Graph) AddVertex(vertex string) bool {
	index := g.IndexOfVertex(vertex)
	// 节点已存在
	if index != -1 {
		return false
	}

	// 超过容量
	if g.VerticeNum >= cap(g.Vertex) {
		g.ensureCapacity()
	}

	g.Vertex[g.VerticeNum] = vertex
	g.VerticeNum++
	return true
}

// ensureCapacity 扩容
func (g *Graph) ensureCapacity() {
	oldCapacity := cap(g.Vertex)
	// 顶点
	growVertex := make([]string, InitVertexNum)
	g.Vertex = append(g.Vertex, growVertex...)

	// 边
	capacity := cap(g.Vertex)
	growEdge := make([][]int, InitVertexNum)
	for i := 0; i < InitVertexNum; i++ {
		growEdge[i] = make([]int, capacity)
	}
	g.Edge = append(g.Edge, growEdge...)
	// 已存在的扩容
	for i := 0; i < oldCapacity; i++ {
		edge := make([]int, InitVertexNum)
		g.Edge[i] = append(g.Edge[i], edge...)
	}
}

// AddEdge 新增边，如果边存在则返回false
// fromVertex 起始顶点
// toVertex 结束顶点
func (g *Graph) AddEdge(fromVertex, toVertex string) bool {
	fromIndex := g.IndexOfVertex(fromVertex)
	if fromIndex == -1 {
		return false
	}
	toIndex := g.IndexOfVertex(toVertex)
	if toIndex == -1 {
		return false
	}

	g.Edge[fromIndex][toIndex] = 1
	g.EdgeNum++
	return true
}

// IsDAG 检测是否为有向无环图
func (g *Graph) IsDAG() bool {
	// 已经是有环图，没必要再次检测
	if !g.isDAG {
		return false
	}

	// 初始化
	capacity := cap(g.Vertex)
	g.color = make([]int, capacity)

	// 对每个顶点遍历
	for i := 0; i < g.VerticeNum; i++ {
		// 该结点后边的顶点都被访问过了，跳过它
		if g.color[i] == -1 {
			continue
		}
		// 从顶点开始进行深度遍历
		g.dfs(i)
		if !g.isDAG {
			return false
		}
	}
	return true
}

// dfs 从index顶点开始进行深度遍历
func (g *Graph) dfs(index int) {
	g.color[index] = 1
	for i := 0; i < g.VerticeNum; i++ {
		// 如果当前顶点有指向的顶点
		if g.Edge[index][i] != 0 {
			// 已经被访问过
			if g.color[i] == 1 {
				// 有环
				g.isDAG = false
				break
			} else if g.color[i] == -1 {
				// 当前顶点后边的顶点都被访问过，直接跳至下一个顶点
				continue
			} else {
				// 递归访问
				g.dfs(i)
			}
		}
	}
	// 遍历过所有相连的顶点后，把本顶点标记为-1
	g.color[index] = -1
}

// NewGraph 创建图
func NewGraph() *Graph {
	edge := make([][]int, InitVertexNum)
	for i := 0; i < InitVertexNum; i++ {
		edge[i] = make([]int, InitVertexNum)
	}
	return &Graph{
		Vertex: make([]string, InitVertexNum),
		Edge:   edge,
		isDAG:  true,
	}
}
