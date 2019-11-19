package main

import (
  "io"
  "os"
  "bufio"
  "fmt"
  "path/filepath"
)

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func GetDBFilePath() string {
  home, err := os.UserHomeDir()
  Check(err)
  return filepath.FromSlash(fmt.Sprintf("%s/%s", home, ".tododb.json"))
}

func WriteDB(data string) error {
  fp := GetDBFilePath()
  f,err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
  Check(err)
  defer f.Close()
  _, err = io.WriteString(f, data+"\n")
  Check(err)
  return f.Sync()
}

func ReadDB() *[]string {
  f,err := os.Open(GetDBFilePath())
  Check(err)
  defer f.Close()
  out := make([]string, 0, 5)
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    out = append(out, scanner.Text())
  }
  return &out
}
