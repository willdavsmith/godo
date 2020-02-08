package main

import (
  "fmt"
  "flag"
  "encoding/json"
)

type Entry struct {
  Description string `json:"description"`
}

func main() {
  priority := flag.Int("p", 0, "priority")
  description := flag.String("d", "Complete task", "description")

  flag.Parse()

  if (contains(flag.Args(), "list")) {
    lines := ReadDB()
    for i,currEntry := range lines {
      tmpEntry := Entry{}
      json.Unmarshal([]byte(currEntry), &tmpEntry)
      if len(lines) > 10 {
        fmt.Printf("[%2d]  %s\n", i, tmpEntry.Description)
      } else {
        fmt.Printf("[%1d]  %s\n", i, tmpEntry.Description)
      }
    }
    fmt.Printf("You have %d items to do\n",len(lines))
  } else {
    lines := ReadDB()
    newEntry := &Entry{Description: *description}
    jsonEntry,_ := json.Marshal(newEntry)

    first := make([]string, *priority)
    copy(first, lines[0:(*priority)])
    second := make([]string, len(lines) - *priority)
    copy(second, lines[(*priority):len(lines)])
    first = append(first, string(jsonEntry))

    lines = append(first, second...)

    fmt.Printf("Added a new task: %s\n", newEntry.Description)
    WriteDB(lines)
  }
}
