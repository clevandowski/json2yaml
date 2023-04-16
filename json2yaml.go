package main

import (
  "encoding/json"
  "fmt"
  "gopkg.in/yaml.v3"
  "io"
  "log"
  "os"
  "strings"
)

func json2yaml(data string) (string, error) {
  data = strings.TrimSpace(data)
  if data == "" {
    return "", nil
  }
  var m any
  err := json.Unmarshal([]byte(data), &m)
  if err != nil {
    return "", err
  }

  output, err := yaml.Marshal(&m)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf("%v", string(output)), nil
}

func main() {
  stdin, err := io.ReadAll(os.Stdin)

  if err != nil {
    panic(err)
  }
  data := string(stdin)

  output, err := json2yaml(data)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  fmt.Printf("%v", string(output))
}
