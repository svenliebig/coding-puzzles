package main

func minJumps(arr []int) int {
  if len(arr) == 1 {
    return 0
  }

  graph := make(map[int][]int)
  for i, v := range arr[1:] {
    if gv, ok := graph[v]; ok {
      graph[v] = append(gv, i + 1)
    } else {
      graph[v] = []int{i + 1}
    }
  }

  queue := make([]int, 1, len(arr))
  queue[0] = 0
  step := 0
  visited := make([]bool, len(arr))
  visited[0] = true
  
  for len(queue) > 0 {
    cl := len(queue)

    for _, v := range queue[:cl] {

      if v == len(arr) - 1 {
        return step
      }

      for _, gv := range graph[arr[v]] {
        if !visited[gv] {
          queue = append(queue, gv)
          visited[gv] = true
          delete(graph, arr[v])
        }
      }

      if v > 0 && !visited[v - 1] {
        queue = append(queue, v - 1)
        visited[v - 1] = true
      }

      if !visited[v + 1] {
        queue = append(queue, v + 1)
        visited[v + 1] = true
      }
    }

    step++
    queue = queue[1:]
  }

  return len(arr)
}
