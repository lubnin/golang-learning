### given graph:
```mermaid
graph LR
    1 --> 5
    2 --> 1
    3 --> 4
    3 --> 2
    4 --> 1
    4 --> 2
    4 --> 5

```

### sorted graph:
```mermaid
graph LR
    1 --> 2
    2 --> 3
    4 --> 5
    3 --> 5
    2 --> 4
    3 --> 4
    1 --> 4

```