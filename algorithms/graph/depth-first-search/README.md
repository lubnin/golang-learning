# Graph Traversal Algorithms

## Introduction

This module contains implementations of two graph traversal algorithms: dfsList and dfsMatrix. These algorithms can be used to explore graphs in a systematic manner.

## dfsList Algorithm

The dfsList algorithm follows the depth-first search (DFS) approach using an adjacency list representation of the graph.

### Algorithm Logic

The algorithm logic can be visualized using the following Mermaid diagram:

```mermaid
graph TD
    A[Start] --> B{Is current vertex visited?}
    B -- Yes --> C[Continue to next vertex]
    B -- No --> D[Print current vertex]
    D --> E[Mark current vertex as visited]
    E --> F{Are there adjacent vertices?}
    F -- Yes --> G[Select an adjacent vertex]
    G --> H{Is the adjacent vertex visited?}
    H -- Yes --> F
    H -- No --> I[Process the adjacent vertex]

```
## dfsMatrix Algorithm

The dfsMatrix algorithm also follows the depth-first search (DFS) approach but uses an adjacency matrix representation of the graph.

### Algorithm Logic

The algorithm logic can be visualized using the following Mermaid diagram:

```mermaid
    graph TD
    A[Start] --> B{Is current vertex visited?}
    B -- Yes --> C[Continue to next vertex]
    B -- No --> D[Print current vertex]
    D --> E[Mark current vertex as visited]
    E --> F{Are there unvisited neighbors?}
    F -- Yes --> G[Select an unvisited neighbor]
    G --> B
    F -- No --> C
    C[End]
```
