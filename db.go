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

func WriteDB(data []string) error {
  fp := GetDBFilePath()
  f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, 0664)
  Check(err)
  defer f.Close()
  for i := 0; i < len(data); i++ {
    _, err = io.WriteString(f, data[i]+"\n")
    Check(err)
  }
  return f.Sync()
}

func ReadDB() []string {
  f, err := os.Open(GetDBFilePath())
  Check(err)
  defer f.Close()
  out := make([]string, 0, 5)
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    out = append(out, scanner.Text())
  }
  return out
}

func ClearDB() {
  fp := GetDBFilePath()
  err := os.Remove(fp)
  Check(err)
  f, e := os.Create(fp)
  Check(e)
  defer f.Close()
}
