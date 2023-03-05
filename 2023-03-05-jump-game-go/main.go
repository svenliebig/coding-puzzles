package main

func minJumps(arr []int) int {
  indexes := make(map[int][]int)

  for i := 1; i < len(arr); i++ {
    if _, ok := indexes[arr[i]]; !ok {
      indexes[arr[i]] = []int{i}
    } else {
      indexes[arr[i]] = append(indexes[arr[i]], i)
    }
  }

  visited := make([]bool, len(arr))
  visited[0] = true
  // a queue of indexes of the array that we will check
  queue := []int{0}
  steps := 0

  for len(queue) > 0 {
    cl := len(queue)
    for _, current := range queue[:cl] {
      if current == len(arr) - 1 {
        return steps
      }

      // check for other indexes that contains this number
      for _, v := range indexes[arr[current]] {
        if !visited[v] {
          queue = append(queue, v)
          visited[v] = true
          delete(indexes, arr[current])
        }
      }

      // add the next, if the next wasn't visited already
      if !visited[current + 1] {
        queue = append(queue, current + 1)
        visited[current + 1] = true
      }

      // add the previous, if the previous wasn't visited already
      if current > 0 && !visited[current - 1] {
        queue = append(queue, current - 1)
        visited[current - 1] = true
      }
    }
    
    queue = queue[cl:]
    steps++
  }

  return len(arr) - 1
}
