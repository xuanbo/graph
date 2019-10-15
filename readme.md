# graph

> 图，邻接矩阵实现

## 功能

- 顶点、边
- 动态扩容（初始化10个顶点）
- DAG（有向无环图）检测

## 使用

```go
package main

import (
    "fmt"

    "github.com/xuanbo/graph"
)

func main() {
    // 创建
    g := graph.NewGraph()
    // 添加顶点
    g.AddVertex("a")
    g.AddVertex("b")
    g.AddVertex("c")
    g.AddVertex("d")
    g.AddVertex("e")
    g.AddVertex("f")
    g.AddVertex("g")
    g.AddVertex("h")
    g.AddVertex("i")
    g.AddVertex("j")
    // 默认初始化10个顶点，此处扩容一次
    g.AddVertex("k")
    // 添加边
    g.AddEdge("a", "b")
    g.AddEdge("b", "c")
    g.AddEdge("c", "d")
    g.AddEdge("d", "a")
    // isDAG
    fmt.Printf("isDAG: %v\n", g.IsDAG())
}
```

注意，底层实现为线程不安全
