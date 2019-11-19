package main

import (
  "fmt"
  "flag"
  "encoding/json"
  "container/heap"
)

func main() {
  priority := flag.Int("p", 1, "priority")
  description := flag.String("d", "", "description")

  flag.Parse()
  if (contains(flag.Args(), "list")) {
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    lines := ReadDB()
    for _,currItem := range *lines {
      tmpItem := Item{}
      json.Unmarshal([]byte(currItem), &tmpItem)
      heap.Push(&pq, &tmpItem)
    }
    for pq.Len() > 0 {
      tmpItem := heap.Pop(&pq).(*Item)
      fmt.Printf("[%2d]  %s\n", tmpItem.Priority, tmpItem.Description)
    }
    fmt.Printf("You have %d items to do\n",len((*lines)))
  } else {
    newItem := &Item{
      Priority: *priority,
      Description: *description}

    jsonItem,_ := json.Marshal(newItem)
    fmt.Printf("Added a new task: %s\n", newItem.Description)
    WriteDB(string(jsonItem))
  }
}
