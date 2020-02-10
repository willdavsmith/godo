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
  var priority int
  var description string
  var remove int
  flag.IntVar(&priority, "p", 0, "priority")
  flag.StringVar(&description, "d", "", "description")
  flag.IntVar(&remove, "r", 0, "remove")

  flag.Parse()
  lines := ReadDB()

  if (contains(flag.Args(), "list")) {
    for i,currEntry := range lines {
      tmpEntry := Entry{}
      json.Unmarshal([]byte(currEntry), &tmpEntry)
      if len(lines) > 10 {
        fmt.Printf("[%2d]  %s\n", i+1, tmpEntry.Description)
      } else {
        fmt.Printf("[%1d]  %s\n", i+1, tmpEntry.Description)
      }
    }
    plural := "s"
    if len(lines) == 1 {
      plural = ""
    }
    fmt.Printf("You have %d item%s to do\n", len(lines), plural)
  } else if priority > 0 {
    if priority > len(lines) {
      priority = len(lines) + 1
    }
    newEntry := &Entry{Description: description}
    jsonEntry,_ := json.Marshal(newEntry)
    first := make([]string, priority-1)
    copy(first, lines[0:priority-1])
    first = append(first, string(jsonEntry))
    second := make([]string, len(lines) - (priority - 1))
    copy(second, lines[priority-1:len(lines)])
    lines = append(first, second...)
    WriteDB(lines)
    fmt.Printf("Added a new task: %s at priority %d\n", newEntry.Description, priority)
  } else if remove > 0 && remove <= len(lines) {
    tmpEntry := Entry{}
    json.Unmarshal([]byte(lines[remove-1]), &tmpEntry)
    ClearDB()
    first := make([]string, remove-1)
    copy(first, lines[0:remove-1])
    second := make([]string, len(lines) - remove)
    copy(second, lines[remove:len(lines)])
    lines = append(first, second...)
    WriteDB(lines)
    fmt.Printf("Completed a task: %s at priority %d\n", tmpEntry.Description, remove)
  }
}
