package main

import
(
  "fmt"
  "github.com/sethvargo/go-password/password"
)

func main() {
  res, err := password.Generate(64, 8, 8, false, false)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(res)
}
