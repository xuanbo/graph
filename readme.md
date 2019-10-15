# graph

> 图，邻接矩阵实现

## 功能

- 顶点、边

- 动态扩容

    初始化10个顶点，每次扩容增加10

- DAG（有向无环图）检测

    基于深度遍历算法

- 遍历DAG

    适合基于DAG的任务流，依次执行的场景

    1. 输出入度为0点顶点
    1. 删除入度为0点顶点的边
    1. 执行1

    优势，不需要起止节点，节点顺序执行，任务等待所有前置节点执行完再执行

    不足，存在并行跨层次等待问题。

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
    g.AddEdge("a", "k")
    g.AddEdge("b", "c")
    g.AddEdge("b", "k")
    g.AddEdge("c", "d")
    g.AddEdge("c", "g")
    g.AddEdge("c", "k")
    // isDAG
    // 输出：false
    fmt.Printf("isDAG: %v\n", g.IsDAG())
    // 遍历DAG
    // 输出：a e f h i j b c d g k
    fmt.Printf("walk: %v\n", g.ForeachDAG())
}
```

注意，底层实现为线程不安全

## 说明

研究基于DAG的任务流系统有感～
